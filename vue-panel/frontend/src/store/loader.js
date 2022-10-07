import client from './httpclient';

export default {
    state: {
        LoaderData: [],
    },
    getters: {
        LOADER_DATA: state => state.LoaderData,
    },
    mutations: {
        SET_LOADER_DATA: (state, payload) => {
            state.LoaderData = payload;
        },
    },
    actions: {
        GET_LOADER_DATA: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/loader/rules/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_LOADER_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        RUN_LOADER_RULE: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/loader/run/',
                data: {'cookie': rootGetters.USER_COOKIE, form},
            }).then(resp => {
                const err = resp.data.err;
                if (err != "") {
                    commit('SET_LAST_ERROR', err);
                }
            })
            .catch(error => {
                commit('SET_LAST_ERROR', error);
                console.error(error);
            });
        },
        CREATE_LOADER_RULE: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/loader/create/',
                data: {'cookie': rootGetters.USER_COOKIE, form},
            }).then(resp => {
                const err = resp.data.err;
                if (err != "") {
                    commit('SET_LAST_ERROR', err);
                }
            })
            .catch(error => {
                commit('SET_LAST_ERROR', error);
                console.error(error);
            });
        },
        EDIT_LOADER_RULE: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/loader/edit/',
                data: {'cookie': rootGetters.USER_COOKIE, form},
            }).then(resp => {
                const err = resp.data.err;
                if (err != "") {
                    commit('SET_LAST_ERROR', err);
                }
            })
            .catch(error => {
                commit('SET_LAST_ERROR', error);
                console.error(error);
            });
        },
        DEL_LOADER_RULE: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/loader/delete/',
                data: {'cookie': rootGetters.USER_COOKIE, form},
            }).then(resp => {
                const err = resp.data.err;
                if (err != "") {
                    commit('SET_LAST_ERROR', err);
                }
            })
            .catch(error => {
                commit('SET_LAST_ERROR', error);
                console.error(error);
            });
        },
    }
}