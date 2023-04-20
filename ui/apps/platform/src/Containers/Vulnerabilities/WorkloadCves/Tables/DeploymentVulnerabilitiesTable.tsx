import React from 'react';
import { Button, ButtonVariant } from '@patternfly/react-core';
import {
    ExpandableRowContent,
    TableComposable,
    Tbody,
    Td,
    Th,
    Thead,
    Tr,
} from '@patternfly/react-table';
import { SVGIconProps } from '@patternfly/react-icons/dist/js/createIcon';
import { gql } from '@apollo/client';

import LinkShim from 'Components/PatternFly/LinkShim';
import SeverityIcons from 'Components/PatternFly/SeverityIcons';
import useSet from 'hooks/useSet';
import { vulnerabilitySeverityLabels } from 'messages/common';
import { getDistanceStrictAsPhrase } from 'utils/dateUtils';
import { UseURLSortResult } from 'hooks/useURLSort';
import { FixableIcon, NotFixableIcon } from 'Components/PatternFly/FixabilityIcons';
import { getEntityPagePath } from '../searchUtils';
import { DynamicColumnIcon } from '../components/DynamicIcon';

import EmptyTableResults from '../components/EmptyTableResults';
import DeploymentComponentVulnerabilitiesTable, {
    DeploymentComponentVulnerability,
    ImageMetadataContext,
    deploymentComponentVulnerabilitiesFragment,
} from './DeploymentComponentVulnerabilitiesTable';

export const deploymentWithVulnerabilitiesFragment = gql`
    ${deploymentComponentVulnerabilitiesFragment}
    fragment DeploymentWithVulnerabilities on Deployment {
        id
        images(query: $query) {
            ...ImageMetadataContext
        }
        imageVulnerabilities(query: $query, pagination: $pagination) {
            id
            cve
            summary
            images(query: $query) {
                imageId: id
                imageComponents(query: $query) {
                    ...DeploymentComponentVulnerabilities
                }
            }
        }
    }
`;

export type DeploymentWithVulnerabilities = {
    id: string;
    images: ImageMetadataContext[];
    imageVulnerabilities: {
        id: string;
        cve: string;
        summary: string;
        images: {
            imageId: string;
            imageComponents: DeploymentComponentVulnerability[];
        }[];
    }[];
};

function formatVulnerabilityData(deployment: DeploymentWithVulnerabilities): {
    id: string;
    cve: string;
    severity: string;
    isFixable: boolean;
    discoveredAtImage: Date | null;
    summary: string;
    images: {
        imageMetadataContext: ImageMetadataContext;
        componentVulnerabilities: DeploymentComponentVulnerability[];
    }[];
}[] {
    const imageMap: Record<string, ImageMetadataContext> = {};
    deployment.images.forEach((image) => {
        imageMap[image.id] = image;
    });

    const vulnerabilities: {
        id: string;
        cve: string;
        severity: string;
        isFixable: boolean;
        discoveredAtImage: Date | null;
        summary: string;
        images: {
            imageMetadataContext: ImageMetadataContext;
            componentVulnerabilities: DeploymentComponentVulnerability[];
        }[];
    }[] = [];

    deployment.imageVulnerabilities.forEach((vulnerability) => {
        const { id, cve, summary, images } = vulnerability;
        // TODO Calculate these
        const severity = 'CRITICAL_VULNERABILITY_SEVERITY';
        const isFixable = false;
        const discoveredAtImage = new Date();

        const vuln = {
            id,
            cve,
            severity,
            isFixable,
            discoveredAtImage,
            summary,
            images: images.map((img) => ({
                imageMetadataContext: imageMap[img.imageId],
                componentVulnerabilities: img.imageComponents,
            })),
        };
        vulnerabilities.push(vuln);
    });

    return vulnerabilities;
}

export type DeploymentVulnerabilitiesTableProps = {
    deployment: DeploymentWithVulnerabilities;
    getSortParams: UseURLSortResult['getSortParams'];
    isFiltered: boolean;
};

function DeploymentVulnerabilitiesTable({
    deployment,
    getSortParams,
    isFiltered,
}: DeploymentVulnerabilitiesTableProps) {
    const expandedRowSet = useSet<string>();

    const vulnerabilities = formatVulnerabilityData(deployment);

    return (
        <TableComposable variant="compact">
            <Thead>
                <Tr>
                    <Th>{/* Header for expanded column */}</Th>
                    <Th sort={getSortParams('CVE')}>CVE</Th>
                    <Th>Severity</Th>
                    <Th>
                        CVE Status
                        {isFiltered && <DynamicColumnIcon />}
                    </Th>
                    <Th>
                        Affected components
                        {isFiltered && <DynamicColumnIcon />}
                    </Th>
                    <Th>First discovered</Th>
                </Tr>
            </Thead>
            {vulnerabilities.length === 0 && <EmptyTableResults colSpan={7} />}
            {vulnerabilities.map(
                ({ cve, severity, summary, isFixable, images, discoveredAtImage }, rowIndex) => {
                    const SeverityIcon: React.FC<SVGIconProps> | undefined =
                        SeverityIcons[severity];
                    const severityLabel: string | undefined = vulnerabilitySeverityLabels[severity];
                    const isExpanded = expandedRowSet.has(cve);

                    const FixabilityIcon = isFixable ? FixableIcon : NotFixableIcon;

                    return (
                        <Tbody key={cve} isExpanded={isExpanded}>
                            <Tr>
                                <Td
                                    expand={{
                                        rowIndex,
                                        isExpanded,
                                        onToggle: () => expandedRowSet.toggle(cve),
                                    }}
                                />
                                <Td dataLabel="CVE">
                                    <Button
                                        variant={ButtonVariant.link}
                                        isInline
                                        component={LinkShim}
                                        href={getEntityPagePath('CVE', cve)}
                                    >
                                        {cve}
                                    </Button>
                                </Td>
                                <Td dataLabel="Severity">
                                    <span>
                                        {SeverityIcon && (
                                            <SeverityIcon className="pf-u-display-inline" />
                                        )}
                                        {severityLabel && (
                                            <span className="pf-u-pl-sm">{severityLabel}</span>
                                        )}
                                    </span>
                                </Td>
                                <Td dataLabel="CVE Status">
                                    <span>
                                        <FixabilityIcon className="pf-u-display-inline" />
                                        <span className="pf-u-pl-sm">
                                            {isFixable ? 'Fixable' : 'Not fixable'}
                                        </span>
                                    </span>
                                </Td>
                                <Td dataLabel="Affected components">
                                    TODO
                                    {
                                        // TODO Distinct component name
                                        /*
                                    {imageComponents.length === 1
                                        ? imageComponents[0].name
                                        : `${imageComponents.length} components`}
                                    */
                                    }
                                </Td>
                                <Td dataLabel="First discovered">
                                    {getDistanceStrictAsPhrase(discoveredAtImage, new Date())}
                                </Td>
                            </Tr>
                            <Tr isExpanded={isExpanded}>
                                <Td />
                                <Td colSpan={6}>
                                    <ExpandableRowContent>
                                        <p className="pf-u-mb-md">{summary}</p>
                                        <DeploymentComponentVulnerabilitiesTable images={images} />
                                    </ExpandableRowContent>
                                </Td>
                            </Tr>
                        </Tbody>
                    );
                }
            )}
        </TableComposable>
    );
}

export default DeploymentVulnerabilitiesTable;
