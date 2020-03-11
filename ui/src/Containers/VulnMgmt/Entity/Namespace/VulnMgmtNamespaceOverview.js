import React, { useContext } from 'react';
import { Link } from 'react-router-dom';
import pluralize from 'pluralize';

import CollapsibleSection from 'Components/CollapsibleSection';
import StatusChip from 'Components/StatusChip';
import RiskScore from 'Components/RiskScore';
import Metadata from 'Components/Metadata';
import Tabs from 'Components/Tabs';
import TabContent from 'Components/TabContent';
import entityTypes from 'constants/entityTypes';
import workflowStateContext from 'Containers/workflowStateContext';
import TopRiskyEntitiesByVulnerabilities from 'Containers/VulnMgmt/widgets/TopRiskyEntitiesByVulnerabilities';
import RecentlyDetectedVulnerabilities from 'Containers/VulnMgmt/widgets/RecentlyDetectedVulnerabilities';
import TopRiskiestImagesAndComponents from 'Containers/VulnMgmt/widgets/TopRiskiestImagesAndComponents';
import DeploymentsWithMostSeverePolicyViolations from 'Containers/VulnMgmt/widgets/DeploymentsWithMostSeverePolicyViolations';
import { getPolicyTableColumns } from 'Containers/VulnMgmt/List/Policies/VulnMgmtListPolicies';
import { entityGridContainerClassName } from 'Containers/Workflow/WorkflowEntityPage';

import RelatedEntitiesSideList from '../RelatedEntitiesSideList';
import TableWidgetFixableCves from '../TableWidgetFixableCves';
import TableWidget from '../TableWidget';

const emptyNamespace = {
    deploymentCount: 0,
    componentCount: 0,
    metadata: {
        clusterName: '',
        clusterId: '',
        name: '',
        priority: 0,
        labels: [],
        id: ''
    },
    policyStatus: {
        status: '',
        failingPolicies: []
    },
    vulnCount: 0,
    vulnerabilities: []
};

const VulnMgmtNamespaceOverview = ({ data, entityContext }) => {
    const workflowState = useContext(workflowStateContext);

    // guard against incomplete GraphQL-cached data
    const safeData = {
        ...emptyNamespace,
        ...data
    };

    const { metadata, policyStatus } = safeData;

    if (!metadata || !policyStatus) return null;

    const { clusterName, clusterId, priority, labels, id } = metadata;
    const { failingPolicies, status } = policyStatus;
    const metadataKeyValuePairs = [];

    if (!entityContext[entityTypes.CLUSTER]) {
        const clusterLink = workflowState.pushRelatedEntity(entityTypes.CLUSTER, clusterId).toUrl();
        metadataKeyValuePairs.push({
            key: 'Cluster',
            value: <Link to={clusterLink}>{clusterName}</Link>
        });
    }

    const namespaceStats = [
        <RiskScore key="risk-score" score={priority} />,
        <React.Fragment key="policy-status">
            <span className="pb-2">Policy status:</span>
            <StatusChip status={status} size="large" />
        </React.Fragment>
    ];

    const currentEntity = { [entityTypes.NAMESPACE]: id };
    const newEntityContext = { ...entityContext, ...currentEntity };

    return (
        <div className="flex h-full">
            <div className="flex flex-col flex-grow min-w-0">
                <CollapsibleSection title="Namespace Summary">
                    <div className={entityGridContainerClassName}>
                        <div className="s-1">
                            <Metadata
                                className="h-full min-w-48 bg-base-100"
                                keyValuePairs={metadataKeyValuePairs}
                                statTiles={namespaceStats}
                                labels={labels}
                                title="Details & Metadata"
                                bgClass
                            />
                        </div>
                        <div className="sx-1 lg:sx-2 sy-1 min-h-55 h-full">
                            <TopRiskyEntitiesByVulnerabilities
                                defaultSelection={entityTypes.DEPLOYMENT}
                                riskEntityTypes={[entityTypes.DEPLOYMENT, entityTypes.IMAGE]}
                                entityContext={currentEntity}
                                small
                            />
                        </div>
                        <div className="s-1">
                            <RecentlyDetectedVulnerabilities entityContext={currentEntity} />
                        </div>
                        <div className="s-1">
                            <TopRiskiestImagesAndComponents entityContext={currentEntity} />
                        </div>
                        <div className="s-1">
                            <DeploymentsWithMostSeverePolicyViolations
                                entityContext={currentEntity}
                            />
                        </div>
                    </div>
                </CollapsibleSection>
                <CollapsibleSection title="Namespace findings">
                    <div className="flex pdf-page pdf-stretch pdf-new shadow rounded relative rounded bg-base-100 mb-4 ml-4 mr-4">
                        <Tabs
                            hasTabSpacing
                            headers={[{ text: 'Policies' }, { text: 'Fixable CVEs' }]}
                        >
                            <TabContent>
                                <TableWidget
                                    header={`${failingPolicies.length} failing ${pluralize(
                                        entityTypes.POLICY,
                                        failingPolicies.length
                                    )} across this namespace`}
                                    entityType={entityTypes.POLICY}
                                    rows={failingPolicies}
                                    noDataText="No failing policies"
                                    className="bg-base-100"
                                    columns={getPolicyTableColumns(workflowState)}
                                    idAttribute="id"
                                />
                            </TabContent>
                            <TabContent>
                                <TableWidgetFixableCves
                                    workflowState={workflowState}
                                    entityContext={entityContext}
                                    entityType={entityTypes.NAMESPACE}
                                    name={safeData?.metadata?.name}
                                    id={safeData?.metadata?.id}
                                />
                            </TabContent>
                        </Tabs>
                    </div>
                </CollapsibleSection>
            </div>

            <RelatedEntitiesSideList
                entityType={entityTypes.NAMESPACE}
                entityContext={newEntityContext}
                data={safeData}
            />
        </div>
    );
};

export default VulnMgmtNamespaceOverview;
