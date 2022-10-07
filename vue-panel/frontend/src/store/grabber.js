import client from './httpclient';

export default {
    state: {
        GrabberData: [],
    },
    getters: {
        GRABBER_DATA: state => state.GrabberData,
    },
    mutations: {
        SET_GRABBER_DATA: (state, payload) => {
            state.GrabberData = payload;
        },
    },
    actions: {
        GET_GRABBER_DATA: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/grabber/rules/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_GRABBER_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        RUN_GRABBER_RULE: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/grabber/run/',
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
        CREATE_GRABBER_RULE: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/grabber/create/',
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
        EDIT_GRABBER_RULE: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/grabber/edit/',
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
        DEL_GRABBER_RULE: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/grabber/delete/',
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