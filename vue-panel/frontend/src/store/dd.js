import client from './httpclient';

export default {
    state: {
        DdData: [],
    },
    getters: {
        DD_DATA: state => state.DdData,
    },
    mutations: {
        SET_DD_DATA: (state, payload) => {
            state.DdData = payload;
        },
    },
    actions: {
        GET_DD_DATA: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/settings/dd/get/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_DD_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        CREATE_DD_RULE: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/settings/dd/create/',
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
        EDIT_DD_RULE: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/settings/dd/edit/',
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
        DEL_DD_RULE: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/settings/dd/delete/',
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