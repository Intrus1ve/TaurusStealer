import client from './httpclient';

export default {
    state: {
        Config: [],
    },
    getters: {
        CONFIG: state => state.Config,
    },
    mutations: {
        SET_CONFIG: (state, payload) => {
            state.Config = payload;
        },
    },
    actions: {
        GET_CONFIG: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/config/get/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_CONFIG', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        UPDATE_CONFIG: async ({commit, rootGetters}, config) => {
            await client({
                method: 'POST',
                url: '/config/update/',
                data: {'cookie': rootGetters.USER_COOKIE, config},
            }).then(resp => {
                const err = resp.data.err;
                if (err != "") {
                    commit('SET_LAST_ERROR', err);
                } else {
                    commit('SET_CONFIG', config);
                }
            })
            .catch(error => {
                commit('SET_LAST_ERROR', error);
                console.error(error);
            });
        },
    }
}