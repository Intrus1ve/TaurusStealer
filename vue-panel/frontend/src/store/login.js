import client from './httpclient';

export default {
    state: {
        User: null,
        Cookie: null,
        IsRoot: null,
    },
    getters: {
        CURRENT_USER: state => {
            if (state.User === null) {
                state.User = localStorage.getItem('User');
            }
            return state.User ;
        },
        USER_COOKIE: state => {
            if (state.Cookie === null) {
                state.Cookie = localStorage.getItem('Cookie');
            }
            return state.Cookie;
        },
        IS_ROOT: state => {
            if (state.IsRoot === null) {
                state.IsRoot = localStorage.getItem('IsRoot');
            }
            return state.IsRoot;
        },
    },
    mutations: {
        LOGIN: (state, payload) => {
            state.User = payload.user;
            state.Cookie = payload.cookie;
            state.IsRoot = payload.isRoot;
            localStorage.setItem('User', payload.user);
            localStorage.setItem('Cookie', payload.cookie);
            localStorage.setItem('IsRoot', payload.isRoot);
        },
        LOGOUT: state => {
            state.User = null;
            state.Cookie = null;
            state.IsRoot = null;
            localStorage.removeItem('User');
            localStorage.removeItem('Cookie');
            localStorage.removeItem('IsRoot');
        },
    },
    actions: {
        LOGIN: async ({commit}, data) => {
            await client({
                method: 'POST',
                url: '/login/',
                data,
            }).then(resp => {
                const user = resp.data['user'];
                const cookie = resp.data['cookie'];
                const isRoot = resp.data['is_root'];
                commit('LOGIN', {user, cookie, isRoot});
            })
            .catch(error => {
                console.error(error);
            });
        },
        LOGOUT: ({commit}) => {
            commit('LOGOUT');
            commit('SET_DAYS_DATA', []);
            commit('SET_LOGS_INFO', []);
            commit('SET_MAP_DATA', []);
            commit('SET_COUNTRIES_DATA', []);
            commit('SET_SOFT_DATA', []);
            commit('SET_PREFIX_DATA', []);
            commit('SET_WIN_DATA', []);
            commit('SET_LOGS_DATA', []);
            commit('SET_LOG_TREE', []);
            commit('SET_FILE_DATA', '');
        },
    }
}