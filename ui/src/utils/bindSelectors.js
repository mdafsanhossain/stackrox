/**
 * Given the map (object) of selectors that work on a particular slice of a global state, and a slicer that extracts this slice form a global state,
 * returns map of selectors with the same keys but with selectors that work on a global state.
 *
 * @template {Object.<string, Function>} T
 * @param {Function} slicer
 * @param {T} selectors
 * @returns {T} map of selectors that can accept as a parameter the same state as slicer accepts
 */
const bindSelectors = (slicer, selectors) =>
    Object.keys(selectors).reduce(
        (boundSelectors, selector) => ({
            ...boundSelectors,
            [selector]: (state, ...args) => selectors[selector](slicer(state), ...args),
        }),
        {}
    );

export default bindSelectors;
