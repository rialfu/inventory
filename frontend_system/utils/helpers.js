export const isAssociativeArray = (val) => {
    return val !== null && typeof val === 'object' && !Array.isArray(val);
};

export const validateFormCustom = (v, schema, state, fields = []) => {
    const result = v.safeParse(schema, state.form);

    if (!result.success) {
        const flattenedErrors = v.flatten(result.issues).nested;
        console.log(v.flatten(result.issues).root || Object.values(flattenedErrors).flat())
        for(const field of fields){
            if(flattenedErrors[field]) state.message[field] = flattenedErrors[field][0]
        }
        state.listMessage = v.flatten(result.issues).root || Object.values(flattenedErrors).flat();
        return false;
    }

    return true; // Validasi berhasil
};