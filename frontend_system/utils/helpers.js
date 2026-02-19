export const isAssociativeArray = (val) => {
    return val !== null && typeof val === 'object' && !Array.isArray(val);
};