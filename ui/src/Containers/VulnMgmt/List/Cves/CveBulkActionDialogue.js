import React, { useState } from 'react';
import PropTypes from 'prop-types';
import gql from 'graphql-tag';
import { useQuery } from 'react-apollo';
import get from 'lodash/get';
import set from 'lodash/set';

import CustomDialogue from 'Components/CustomDialogue';
import InfoList from 'Components/InfoList';
import Loader from 'Components/Loader';
import Message from 'Components/Message';
import { POLICY_ENTITY_ALL_FIELDS_FRAGMENT } from 'Containers/VulnMgmt/VulnMgmt.fragments';
import queryService from 'modules/queryService';
import { createPolicy, savePolicy } from 'services/PoliciesService';
import { truncate } from 'utils/textUtils';

import CveToPolicyShortForm from './CveToPolicyShortForm';

const emptyPolicy = {
    name: '',
    severity: '',
    lifecycleStages: [],
    description: '',
    disabled: false,
    categories: ['Vulnerability Management'],
    fields: {
        cve: ''
    },
    whitelists: []
};

const CveBulkActionDialogue = ({ closeAction, bulkActionCveIds }) => {
    const [messageObj, setMessageObj] = useState(null);

    // the combined CVEs are used for both the policy object and the GraphQL query var
    const cvesStr = bulkActionCveIds.join(',');

    // prepare policy object
    const [policyIdentifer, setPolicyIdentifier] = useState('');

    // prepare policy object
    const populatedPolicy = { ...emptyPolicy, fields: { cve: cvesStr } };
    const [policy, setPolicy] = useState(populatedPolicy);

    // use GraphQL to get the (hopefully cached) cve summaries to display in the dialog
    const CVES_QUERY = gql`
        query getCves($query: String) {
            results: vulnerabilities(query: $query) {
                id: cve
                cve
                summary
            }
        }
    `;
    const cvesObj = {
        cve: cvesStr
    };
    const cveQueryOptions = {
        variables: {
            query: queryService.objectToWhereClause(cvesObj)
        }
    };
    const { loading: cveLoading, data: cveData } = useQuery(CVES_QUERY, cveQueryOptions);
    const cveItems =
        !cveLoading && cveData && cveData.results && cveData.results.length ? cveData.results : [];

    // use GraphQL to get existing vulnerability-related policies
    const POLICIES_QUERY = gql`
        query getPolicies($policyQuery: String) {
            results: policies(query: $policyQuery) {
                ...policyFields
            }
        }
        ${POLICY_ENTITY_ALL_FIELDS_FRAGMENT}
    `;
    const policyQueryOptions = {
        variables: {
            policyQuery: queryService.objectToWhereClause({
                Category: 'Vulnerability Management'
            })
        }
    };
    const { loading: policyLoading, data: policyData } = useQuery(
        POLICIES_QUERY,
        policyQueryOptions
    );
    const existingPolicies =
        !policyLoading && policyData && policyData.results && policyData.results.length
            ? policyData.results.map(pol => ({ ...pol, value: pol.id, label: pol.name }))
            : [];

    function handleChange(event) {
        if (get(policy, event.target.name) !== undefined) {
            const newPolicyFields = { ...policy };
            const newValue =
                event.target.type === 'checkbox' ? event.target.checked : event.target.value;
            set(newPolicyFields, event.target.name, newValue);
            setPolicy(newPolicyFields);
        }
    }

    function setSelectedPolicy(value) {
        // some string was typed or selected
        const existingPolicy =
            existingPolicies && existingPolicies.find(pol => pol.value === value);

        if (existingPolicy) {
            // it matches an existing policy's ID, so must have been selected from existing list
            const newCveList = existingPolicy.fields.cve
                ? `${existingPolicy.fields.cve},${cvesStr}`
                : cvesStr;
            const newFields = { ...existingPolicy.fields, cve: newCveList };

            const newPolicy = { ...existingPolicy, fields: newFields, id: value };

            setPolicy(newPolicy);
        } else {
            // not in existing list, so must be a typed name instead of an ID
            const newPolicy = { ...policy, name: value, id: '' };

            setPolicy(newPolicy);
        }

        // update the form
        setPolicyIdentifier(value);
    }

    function handleClose(idsToStaySelected) {
        closeAction(idsToStaySelected);
    }

    function closeWithoutSaving() {
        handleClose(bulkActionCveIds);
    }

    function addToPolicy() {
        // TODO: make the form submission more robust
        //   this current save function is only for smoke-testing the form
        const addToFunc = policy.id ? savePolicy : createPolicy;

        addToFunc(policy)
            .then(() => {
                setMessageObj({ type: 'info', message: 'Policy successfully saved' });

                // close the dialog after giving the user a little time to process the success message
                setTimeout(handleClose, 3000);
            })
            .catch(error => {
                setMessageObj({
                    type: 'error',
                    message: `Policy could not be saved. Please try again. (${error})`
                });

                // hide the error message after giving the user time to read it
                setTimeout(() => {
                    setMessageObj(null);
                }, 7000);
            });
    }

    function renderCve(item) {
        const truncatedSummary = truncate(item.summary, 120);
        return (
            <li key={item.id} className="flex items-center bg-tertiary-200 mb-2 p-2">
                <span className="min-w-32">{item.cve}</span>
                <span>{truncatedSummary}</span>
            </li>
        );
    }

    // render section
    if (bulkActionCveIds.length === 0) return null;

    return (
        <CustomDialogue
            className="max-w-3/4 md:max-w-2/3 lg:max-w-1/2"
            title="Add To Policy"
            text=""
            onConfirm={addToPolicy}
            confirmText="Save Policy"
            confirmDisabled={messageObj}
            onCancel={closeWithoutSaving}
        >
            {/* TODO: replace with working form, this is a temporary placeholder only */}
            <div className="p-4">
                {messageObj && <Message type={messageObj.type} message={messageObj.message} />}
                <CveToPolicyShortForm
                    policy={policy}
                    handleChange={handleChange}
                    existingPolicies={existingPolicies}
                    selectedPolicy={policyIdentifer}
                    setSelectedPolicy={setSelectedPolicy}
                />
                <div className="pt-2">
                    <h3 className="mb-2">{`${
                        bulkActionCveIds.length
                    } CVEs listed below will be added to this policy:`}</h3>
                    {cveLoading && <Loader />}
                    {!cveLoading && (
                        <InfoList items={cveItems} renderItem={renderCve} extraClassNames="h-48" />
                    )}
                </div>
            </div>
        </CustomDialogue>
    );
};

CveBulkActionDialogue.propTypes = {
    closeAction: PropTypes.func.isRequired,
    bulkActionCveIds: PropTypes.arrayOf(PropTypes.string).isRequired
};

export default CveBulkActionDialogue;
