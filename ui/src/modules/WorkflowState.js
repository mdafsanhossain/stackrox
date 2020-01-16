import cloneDeep from 'lodash/cloneDeep';
import uniqBy from 'lodash/uniqBy';

import entityRelationships from 'modules/entityRelationships';
import generateURL from 'modules/URLGenerator';
import { searchParams, sortParams, pagingParams } from 'constants/searchParams';

// An item in the workflow stack
export class WorkflowEntity {
    constructor(entityType, entityId) {
        if (entityType) {
            this.t = entityType;
        }
        if (entityId) {
            this.i = entityId;
        }
        Object.freeze(this);
    }

    get entityType() {
        return this.t;
    }

    get entityId() {
        return this.i;
    }
}

// Returns true if stack provided makes sense
export function isStackValid(stack) {
    const existingEntityTypes = uniqBy(stack, 'entityType');

    // if the stack is smaller or equal to two entity types, it is always valid
    if (Object.keys(existingEntityTypes).length <= 2) return true;

    // stack is invalid when the stack is in one of three states:
    //
    // 1) entity -> (entity parent list) -> entity parent -> nav away
    // 2) entity -> (entity matches list) -> match entity -> nav away
    // 3) entity -> ... -> same entity (nav away)

    let isParentState;
    let isMatchState;
    let isDuplicateState;

    const entityTypeMap = {};

    stack.forEach((entity, i) => {
        const { entityType, entityId } = entity;

        if (entityTypeMap[entityType]) {
            isDuplicateState = !!entityTypeMap[entityType];
        }
        if (i > 0) {
            entityTypeMap[entityType] = entityId;
            const { entityType: prevType } = stack[i - 1];
            if (prevType !== entityType) {
                if (!isParentState) {
                    // this checks if the current type on the stack is a parent of the previous type
                    const isParent = entityRelationships.isContained(entityType, prevType);
                    isParentState = i !== stack.length - 1 && isParent;
                }
                if (!isMatchState) {
                    const isContained = entityRelationships.isContained(prevType, entityType);
                    // if prev entity type contains current entity type, match state doesn't matter and stack is valid
                    if (!isContained) {
                        // extended matches navigate away
                        const isExtendedMatch = entityRelationships.isExtendedMatch(
                            prevType,
                            entityType
                        );
                        const upMatch = entityRelationships.isPureMatch(prevType, entityType);
                        const downMatch = entityRelationships.isPureMatch(entityType, prevType);
                        // reflexive matches navigate away if it's not the last relationship on stack
                        const isReflexiveMatchState =
                            i !== stack.length - 1 && upMatch && downMatch;
                        isMatchState = isReflexiveMatchState || isExtendedMatch;
                    }
                }
            }
        }
        return false;
    });
    return !isParentState && !isMatchState && !isDuplicateState;
}

// Resets the current state based on minimal parameters
function baseStateStack(entityType, entityId) {
    return [new WorkflowEntity(entityType, entityId)];
}

// Returns skimmed stack for stack to navigate away to
function skimStack(stack) {
    if (stack.length < 2) return stack;

    const currentItem = stack.slice(-1)[0];
    // if the last item on the stack is an entity, return the entity
    if (currentItem.entityId) return [currentItem];
    // else the last item on the stack is a list, return the previous entity + related list
    return stack.slice(-2);
}

// Checks state stack for overflow state/invalid state and returns a valid skimmed version
function trimStack(stack) {
    // Navigate away if:
    // If there's no more "room" in the stack
    return isStackValid(stack) ? stack : skimStack(stack);
}

/**
 * Summary: Class that ensures the shape of a WorkflowState object
 * {
 *   useCase: 'text',
 *   stateStack: [{t: 'entityType', i: 'entityId'},{t: 'entityType', i: 'entityId'}]
 * }
 */
export class WorkflowState {
    constructor(useCase, stateStack, search, sort, paging) {
        this.useCase = useCase;
        this.stateStack = cloneDeep(stateStack) || [];
        this.search = cloneDeep(search) || {};
        this.sort = cloneDeep(sort) || {};
        this.paging = cloneDeep(paging) || {};

        this.sidePanelActive = this.getPageStack().length !== this.stateStack.length;

        Object.freeze(this);
        Object.freeze(this.search);
        Object.freeze(this.stateStack);
        Object.freeze(this.sort);
        Object.freeze(this.paging);
    }

    clone() {
        const { useCase, stateStack, search, sort, paging } = this;
        return new WorkflowState(useCase, stateStack, search, sort, paging);
    }

    // Returns current entity (top of stack)
    getCurrentEntity() {
        if (!this.stateStack.length) return null;
        return this.stateStack.slice(-1)[0];
    }

    // Returns type of the current entity (top of stack)
    getCurrentEntityType() {
        const currentEntity = this.getCurrentEntity();

        if (!currentEntity) return null;

        return currentEntity.t;
    }

    // Returns base (first) entity of stack
    getBaseEntity() {
        if (!this.stateStack.length) return null;
        return this.stateStack[0];
    }

    // Returns workflow entities related to page level
    getPageStack() {
        const { stateStack } = this;
        if (stateStack.length < 2) return stateStack;

        // list page or entity page with entity sidepanel
        if (!stateStack[0].entityId || (stateStack.length > 1 && stateStack[1].entityId))
            return stateStack.slice(0, 1);

        // entity page with tab
        return stateStack.slice(0, 2);
    }

    // Gets selected table row (first side panel entity)
    getSelectedTableRow() {
        if (this.stateStack.length < 2 || !this.sidePanelActive) return null;
        return this.stateStack.slice(1, 2)[0];
    }

    getCurrentSearchState() {
        const param = this.sidePanelActive ? searchParams.sidePanel : searchParams.page;
        return this.search[param] || {};
    }

    getCurrentSortState() {
        const param = this.sidePanelActive ? sortParams.sidePanel : sortParams.page;
        return this.sort[param] || {};
    }

    getCurrentPagingState() {
        const param = this.sidePanelActive ? pagingParams.sidePanel : pagingParams.page;
        return this.paging[param] || {};
    }

    // Returns skimmed stack version of WorkflowState to render into URL
    getSkimmedStack() {
        const { useCase, stateStack, search, sort, paging } = this;
        const newStateStack = skimStack(stateStack);
        return new WorkflowState(useCase, newStateStack, search, sort, paging);
    }

    // Resets the current state based on minimal parameters
    reset(useCase, entityType, entityId, search, sort, paging) {
        const newUseCase = useCase || this.useCase;
        const newStateStack = baseStateStack(entityType, entityId);
        return new WorkflowState(newUseCase, newStateStack, search, sort, paging);
    }

    resetPage(type, id) {
        const newStateStack = [new WorkflowEntity(type, id)];

        const { useCase } = this;
        return new WorkflowState(useCase, newStateStack);
    }

    // Returns a cleared stack on current use case. Useful when building state from scratch.
    clear() {
        const newStateStack = [];
        const { useCase } = this;
        return new WorkflowState(useCase, newStateStack);
    }

    // sets the stateStack to base state when returning from side panel
    removeSidePanelParams() {
        const { useCase, search, sort, paging } = this;
        const newStateStack = this.getPageStack();
        const newSearch = search ? { [searchParams.page]: search[searchParams.page] } : null;
        const newSort = sort ? { [searchParams.page]: sort[searchParams.page] } : null;
        const newPaging = paging ? { [searchParams.page]: paging[searchParams.page] } : null;
        return new WorkflowState(useCase, newStateStack, newSearch, newSort, newPaging);
    }

    // sets statestack to only the first item
    base() {
        const { useCase, stateStack } = this;
        return new WorkflowState(useCase, stateStack.slice(0, 1));
    }

    // Adds a list of entityType related to the current workflowState
    pushList(type) {
        const { useCase, stateStack, search, sort, paging } = this;
        const newItem = new WorkflowEntity(type);
        const currentItem = this.getCurrentEntity();

        // Slice an item off the end of the stack if this push should result in a replacement (e.g. clicking on tabs)
        const newStateStack =
            currentItem && currentItem.entityType && !currentItem.entityId
                ? stateStack.slice(0, -1)
                : [...stateStack];
        newStateStack.push(newItem);

        return new WorkflowState(useCase, trimStack(newStateStack), search, sort, paging);
    }

    // Selects an item in a list by Id
    pushListItem(id) {
        const { useCase, stateStack, search, sort, paging } = this;
        const currentItem = this.getCurrentEntity();
        const newItem = new WorkflowEntity(currentItem.entityType, id);
        // Slice an item off the end of the stack if this push should result in a replacement (e.g. clicking on multiple list items)
        const newStateStack = currentItem.entityId ? stateStack.slice(0, -1) : [...stateStack];
        newStateStack.push(newItem);

        return new WorkflowState(useCase, newStateStack, search, sort, paging);
    }

    // Shows an entity in relation to the top entity in the workflow
    pushRelatedEntity(type, id) {
        const { useCase, stateStack, search, sort, paging } = this;
        const currentItem = stateStack.slice(-1)[0];

        if (currentItem && !currentItem.entityId) return this;

        const newStateStack = trimStack([...stateStack, new WorkflowEntity(type, id)]);

        return new WorkflowState(useCase, newStateStack, search, sort, paging);
    }

    // Goes back one level to the nearest valid state
    pop() {
        if (this.stateStack.length === 1)
            // A state stack has to have at least one item in it
            return this;

        const { useCase, stateStack, search, sort, paging } = this;

        return new WorkflowState(
            useCase,
            stateStack.slice(0, stateStack.length - 1),
            search,
            sort,
            paging
        );
    }

    setSearch(newProps) {
        const { useCase, stateStack, search, sort, paging, sidePanelActive } = this;
        const param = sidePanelActive ? searchParams.sidePanel : searchParams.page;

        const newSearch = {
            ...search,
            [param]: newProps
        };
        return new WorkflowState(useCase, stateStack, newSearch, sort, paging);
    }

    setSort(sortProp) {
        const { useCase, stateStack, search, sort, paging, sidePanelActive } = this;
        const param = sidePanelActive ? sortParams.sidePanel : sortParams.page;

        const newSort = {
            ...sort,
            [param]: sortProp
        };

        return new WorkflowState(useCase, stateStack, search, newSort, paging);
    }

    clearSort() {
        const { useCase, stateStack, search, sort, paging, sidePanelActive } = this;
        const param = sidePanelActive ? sortParams.sidePanel : sortParams.page;

        const newSort = {
            ...sort,
            [param]: undefined
        };

        return new WorkflowState(useCase, stateStack, search, newSort, paging);
    }

    setPage(pagingProp) {
        const { useCase, stateStack, search, sort, paging, sidePanelActive } = this;
        const param = sidePanelActive ? pagingParams.sidePanel : pagingParams.page;

        const newPaging = {
            ...paging,
            [param]: pagingProp
        };
        return new WorkflowState(useCase, stateStack, search, sort, newPaging);
    }

    toUrl() {
        return generateURL(this);
    }

    getEntityContext() {
        return this.stateStack
            .filter(item => !!item.entityId)
            .reduce((entityContext, item) => {
                return { ...entityContext, [item.entityType]: item.entityId };
            }, {});
    }
}
