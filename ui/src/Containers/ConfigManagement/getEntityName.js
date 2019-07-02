import entityTypes from 'constants/entityTypes';
import resolvePath from 'object-resolve-path';
import isEmpty from 'lodash/isEmpty';

const entityNameKeyMap = {
    [entityTypes.SERVICE_ACCOUNT]: data => resolvePath(data, 'serviceAccount.name'),
    [entityTypes.SECRET]: data => resolvePath(data, 'secret.name'),
    [entityTypes.CLUSTER]: data => resolvePath(data, 'results.name'),
    [entityTypes.DEPLOYMENT]: data => resolvePath(data, 'deployment.name'),
    [entityTypes.NAMESPACE]: data => resolvePath(data, 'results.metadata.name'),
    [entityTypes.ROLE]: data => resolvePath(data, 'clusters[0].k8srole.name'),
    [entityTypes.NODE]: data => resolvePath(data, 'node.name'),
    [entityTypes.CONTROL]: data => {
        if (!data.results) return null;
        return `${data.results.name} - ${data.results.description}`;
    },
    [entityTypes.IMAGE]: data => resolvePath(data, 'image.name.fullName'),
    [entityTypes.POLICY]: data => resolvePath(data, 'policy.name')
};

const getEntityName = (entityType, data) => {
    if (isEmpty(data)) return null;
    try {
        return entityNameKeyMap[entityType](data);
    } catch (error) {
        throw new Error(
            `Entity (${entityType}) is not mapped correctly in the "entityToNameResolverMapping"`
        );
    }
};

export default getEntityName;
