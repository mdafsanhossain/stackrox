import React from 'react';
import PropTypes from 'prop-types';
import { Field } from 'redux-form';
import Select, { defaultSelectStyles } from 'Components/ReactSelect';

const ReduxSelect = ({
    input: { name, value, onChange },
    options,
    placeholder,
    disabled,
    customComponents,
    styles,
    defaultValue
}) => (
    <Select
        key={name}
        onChange={onChange}
        options={options}
        placeholder={placeholder}
        value={value || defaultValue}
        isDisabled={disabled}
        components={customComponents}
        styles={styles}
        menuPlacement="auto"
    />
);

ReduxSelect.propTypes = {
    input: PropTypes.shape({
        value: PropTypes.oneOfType([PropTypes.string, PropTypes.bool]),
        name: PropTypes.string,
        onChange: PropTypes.func
    }).isRequired,
    options: PropTypes.arrayOf(PropTypes.shape({})).isRequired,
    placeholder: PropTypes.string.isRequired,
    disabled: PropTypes.bool,
    customComponents: PropTypes.shape({}),
    styles: PropTypes.shape({}),
    defaultValue: PropTypes.string
};

ReduxSelect.defaultProps = {
    disabled: false,
    customComponents: {},
    styles: defaultSelectStyles,
    defaultValue: null
};

const ReduxSelectField = ({
    name,
    options,
    placeholder,
    disabled,
    customComponents,
    styles,
    value
}) => (
    <Field
        key={name}
        name={name}
        options={options}
        customComponents={customComponents}
        component={ReduxSelect}
        placeholder={placeholder}
        disabled={disabled}
        styles={styles}
        defaultValue={value}
        className="border bg-base-100 border-base-300 text-base-600 p-3 pr-8 rounded-r-sm cursor-pointer z-50 focus:border-base-300 w-full font-400"
    />
);

ReduxSelectField.propTypes = {
    name: PropTypes.oneOfType([PropTypes.string, PropTypes.bool]).isRequired,
    options: PropTypes.arrayOf(PropTypes.shape({})).isRequired,
    placeholder: PropTypes.string,
    disabled: PropTypes.bool,
    customComponents: PropTypes.shape({}),
    styles: PropTypes.shape({}),
    value: PropTypes.string
};

ReduxSelectField.defaultProps = {
    placeholder: 'Select one...',
    disabled: false,
    customComponents: {},
    styles: defaultSelectStyles,
    value: null
};

export default ReduxSelectField;
