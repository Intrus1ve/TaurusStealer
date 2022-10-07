import client from './httpclient';

export default {
    state: {
        UsersData: [],
    },
    getters: {
        USERS_DATA: state => state.UsersData,
    },
    mutations: {
        SET_USERS_DATA: (state, payload) => {
            state.UsersData = payload;
        },
    },
    actions: {
        GET_USERS_DATA: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/settings/users/get/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_USERS_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        CREATE_USER: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/settings/users/create/',
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
        EDIT_USER: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/settings/users/edit/',
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
        DEL_USER: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/settings/users/delete/',
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