import React from 'react';
import { gql } from '@apollo/client';
import { Flex } from '@patternfly/react-core';
import { VulnerabilitySeverity } from 'types/cve.proto';
import AffectedImages from '../SummaryCards/AffectedImages';
import TopCvssScoreBreakdown from '../SummaryCards/TopCvssScoreBreakdown';
import BySeveritySummaryCard from '../SummaryCards/BySeveritySummaryCard';

export type ImageCveSummaryCount = {
    totalImageCount: number;
};

export type ImageCveSeveritySummary = {
    affectedImageCountBySeverity: {
        critical: number;
        important: number;
        moderate: number;
        low: number;
    };
    affectedImageCount: number;
    topCVSS: number;
};

export const imageCveSeveritySummaryFragment = gql`
    fragment ImageCVESeveritySummary on ImageCVECore {
        affectedImageCountBySeverity {
            critical
            important
            moderate
            low
        }
        affectedImageCount
        topCVSS
        # TODO vector
    }
`;

export const imageCveSummaryCountFragment = gql`
    fragment ImageCVESummaryCounts on Query {
        totalImageCount: imageCount
    }
`;

const defaultSeveritySummary = {
    affectedImageCountBySeverity: {
        critical: 0,
        important: 0,
        moderate: 0,
        low: 0,
    },
    affectedImageCount: 0,
    topCVSS: 0,
};

export type ImageCveSummaryCardsProps = {
    summaryCounts: ImageCveSummaryCount;
    severitySummary: ImageCveSeveritySummary | undefined;
    hiddenSeverities: Set<VulnerabilitySeverity>;
};

function ImageCveSummaryCards({
    summaryCounts,
    severitySummary = defaultSeveritySummary,
    hiddenSeverities,
}: ImageCveSummaryCardsProps) {
    const { affectedImageCount, topCVSS } = severitySummary;
    const { totalImageCount } = summaryCounts;
    return (
        <Flex
            direction={{ default: 'column', lg: 'row' }}
            alignItems={{ lg: 'alignItemsStretch' }}
            justifyContent={{ default: 'justifyContentSpaceBetween' }}
        >
            <AffectedImages
                className="pf-u-flex-grow-1 pf-u-flex-basis-0"
                affectedImageCount={affectedImageCount}
                totalImagesCount={totalImageCount}
            />
            <BySeveritySummaryCard
                className="pf-u-flex-grow-1 pf-u-flex-basis-0"
                title="Images by severity"
                severityCounts={severitySummary.affectedImageCountBySeverity}
                hiddenSeverities={hiddenSeverities}
            />
            <TopCvssScoreBreakdown
                className="pf-u-flex-grow-1 pf-u-flex-basis-0"
                cvssScore={topCVSS}
                vector="TODO - Not implemented"
            />
        </Flex>
    );
}

export default ImageCveSummaryCards;
