/** @type {import('./$types').LayoutLoad} */
export const load = async ({ url: { pathname } }: any) => {
    return { pathname };
};