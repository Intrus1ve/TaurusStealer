import client from './httpclient';

export default {
    state: {
        LogsData: [],
        LogTree: [],
        FileData: '',
    },
    getters: {
        LOGS_DATA: state => state.LogsData,
        LOG_TREE: state => state.LogTree,
        FILE_DATA: state => state.FileData,
    },
    mutations: {
        SET_LOGS_DATA: (state, payload) => {
            state.LogsData = payload;
        },
        SET_LOG_TREE: (state, payload) => {
            state.LogTree = payload;
        },
        SET_FILE_DATA: (state, payload) => {
            state.FileData = payload;
        },
    },
    actions: {
        GET_LOGS_DATA: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/logs/data/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_LOGS_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        SET_LOG_COMMENT: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/logs/comment/',
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
        GET_LOG_TREE: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/logs/logTree/',
                data: {'cookie': rootGetters.USER_COOKIE, form},
            }).then(resp => {
                commit('SET_LOG_TREE', resp.data);
            })
            .catch(error => {
                commit('SET_LAST_ERROR', error);
                console.error(error);
            });
        },
        GET_FILE_DATA: async ({commit, rootGetters}, fileName) => {
            await client({
                method: 'POST',
                url: '/logs/fileData/',
                data: {'cookie': rootGetters.USER_COOKIE, fileName},
            }).then(resp => {
                commit('SET_FILE_DATA', resp.data);
            })
            .catch(error => {
                commit('SET_LAST_ERROR', error);
                console.error(error);
            });
        },
        DOWNLOAD_LOG: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/logs/download/',
                data: {'cookie': rootGetters.USER_COOKIE, form},
                responseType: 'blob'
            }).then(resp => {
                const url = window.URL.createObjectURL(new Blob([resp.data]));
                const link = document.createElement('a');
                link.href = url;
                link.setAttribute('download', form.country+'_'+form.uid+'.zip');
                document.body.appendChild(link);
                link.click();
            })
            .catch(error => {
                commit('SET_LAST_ERROR', error);
                console.error(error);
            });
        },
        DEL_LOG: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/logs/delete/',
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
        FILTER_LOGS: async ({commit, rootGetters}, filter) => {
            await client({
                method: 'POST',
                url: '/logs/filter/',
                data: {'cookie': rootGetters.USER_COOKIE, filter},
            }).then(resp => {
                commit('SET_LOGS_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        SELECTED_LOGS_ACTION: async ({commit, rootGetters}, selected) => {
            await client({
                method: 'POST',
                url: '/logs/selected/',
                data: {'cookie': rootGetters.USER_COOKIE, selected},
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