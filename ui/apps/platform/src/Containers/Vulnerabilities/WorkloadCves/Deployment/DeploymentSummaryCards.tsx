import React from 'react';
import { gql } from '@apollo/client';
import { Grid, GridItem } from '@patternfly/react-core';

import { VulnerabilitySeverity } from 'types/cve.proto';
import CvesByStatusSummaryCard, {
    ImageVulnerabilityCounter,
} from '../SummaryCards/CvesByStatusSummaryCard';
import BySeveritySummaryCard from '../SummaryCards/BySeveritySummaryCard';
import { FixableStatus } from '../types';

export type DeploymentImageCveCountBySeverity = {
    critical: number;
    important: number;
    moderate: number;
    low: number;
};

export const deploymentImageCveCountBySeverityFragment = gql`
    fragment DeploymentImageCveCountBySeverity on ResourceCountByCVESeverity {
        critical
        important
        moderate
        low
    }
`;

export type DeploymentSummaryCardsProps = {
    severityCounts: DeploymentImageCveCountBySeverity;
    statusCounts: ImageVulnerabilityCounter;
    hiddenSeverities: Set<VulnerabilitySeverity>;
    hiddenStatuses: Set<FixableStatus>;
};

function DeploymentSummaryCards({
    severityCounts,
    statusCounts,
    hiddenSeverities,
    hiddenStatuses,
}: DeploymentSummaryCardsProps) {
    return (
        <Grid hasGutter>
            <GridItem sm={12} md={6} xl2={4}>
                <BySeveritySummaryCard
                    title="CVEs by severity"
                    severityCounts={severityCounts}
                    hiddenSeverities={hiddenSeverities}
                />
            </GridItem>
            <GridItem sm={12} md={6} xl2={4}>
                <CvesByStatusSummaryCard
                    cveStatusCounts={statusCounts}
                    hiddenStatuses={hiddenStatuses}
                />
            </GridItem>
        </Grid>
    );
}

export default DeploymentSummaryCards;
