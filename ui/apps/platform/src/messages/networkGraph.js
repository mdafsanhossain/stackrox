import { PROTOCOLS } from 'constants/networkGraph';

const networkProtocolLabels = {
    [PROTOCOLS.L4_PROTOCOL_TCP]: 'TCP',
    [PROTOCOLS.L4_PROTOCOL_UDP]: 'UDP',
    [PROTOCOLS.L4_PROTOCOL_ANY]: 'Any Protocol',
};

export default networkProtocolLabels;
