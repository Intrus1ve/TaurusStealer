import client from './httpclient';

export default {
    state: {
        BackupsData: [],
    },
    getters: {
        BACKUPS_DATA: state => state.BackupsData,
    },
    mutations: {
        SET_BACKUPS_DATA: (state, payload) => {
            state.BackupsData = payload;
        },
    },
    actions: {
        GET_BACKUPS_DATA: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/settings/backups/get/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_BACKUPS_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        CREATE_BACKUP: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/settings/backups/create/',
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
        DEL_BACKUP: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/settings/backups/delete/',
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
        DOWNLOAD_BACKUP: async ({commit, rootGetters}, form) => {
            await client({
                method: 'POST',
                url: '/settings/backups/download/',
                data: {'cookie': rootGetters.USER_COOKIE, form},
                responseType: 'blob',
            }).then(resp => {
                const url = window.URL.createObjectURL(new Blob([resp.data]));
                const link = document.createElement('a');
                link.href = url;
                link.setAttribute('download', 'Taurus_'+form.date+'.zip');
                document.body.appendChild(link);
                link.click();
            })
            .catch(error => {
                commit('SET_LAST_ERROR', error);
                console.error(error);
            });
        },
    }
}