export const baseURL = '/main/vulnerability-management';

export const url = {
    dashboard: baseURL,
    list: {
        policies: `${baseURL}/policies`,
        clusters: `${baseURL}/clusters`,
        namespaces: `${baseURL}/namespaces`,
        deployments: `${baseURL}/deployments`,
        images: `${baseURL}/images`,
        components: `${baseURL}/components`,
        cves: `${baseURL}/cves`
    }
};

export const listSelectors = {
    tableHeader: '.rt-thead.-header',
    tableColumn: '.rt-th.leading-normal > div',
    tableBodyRows: '.rt-tbody > .rt-tr-group > .rt-tr',
    tableBodyColumns: '.rt-tbody > .rt-tr-group > .rt-tr > .rt-td.leading-normal',
    tableColumnLinks: '.rt-tr-group > .rt-tr > .rt-td > a',
    tableCVEColumnLinks: '.rt-tr-group > .rt-tr > .rt-td > .items-center'
};

export const dashboardSelectors = {
    widgets: "[data-test-id='widget']",
    tileLinks: "[data-test-id='tile-link']",
    tileLinkValue: "[data-test-id='tile-link-value']",
    applicationAndInfrastructureDropdown: 'button:contains("Application & Infrastructure")',
    topRiskyItems: {
        widget: '[data-test-id="widget"]',
        select: {
            input: '[data-test-id="widget"] .react-select__control',
            value: '[data-test-id="widget"] .react-select__single-value',
            options: '[data-test-id="widget"] .react-select__option'
        }
    },
    getMenuListItem: name => {
        return `[data-test-id="menu-list"] [data-test-id="${name}"]`;
    },
    getWidget: title => {
        return `[data-test-id="widget"]:contains('${title}')`;
    },
    viewAllButton: 'button:contains("View All")'
};
export const selectors = {
    ...dashboardSelectors,
    ...listSelectors
};
