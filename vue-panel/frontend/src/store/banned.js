import client from './httpclient';

export default {
    state: {
        BannedData: [],
    },
    getters: {
        BANNED_DATA: state => state.BannedData,
    },
    mutations: {
        SET_BANNED_DATA: (state, payload) => {
            state.BannedData = payload;
        },
    },
    actions: {
        GET_BANNED_DATA: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/settings/banned/get/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_BANNED_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        CREATE_BAN_RULE: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/settings/banned/create/',
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
        EDIT_BAN_RULE: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/settings/banned/edit/',
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
        DEL_BAN_RULE: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/settings/banned/delete/',
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