import client from './httpclient';

export default {
    state: {
        DaysData: [],
        LogsInfo: [],
        MapData: [],
        TopCountriesData: [],
        CountriesData: [],
        SoftData: [],
        TopPrefixData: [],
        PrefixData: [],
        WinData: [],
    },
    getters: {
        DAYS_DATA: state => state.DaysData,
        LOGS_INFO: state => state.LogsInfo,
        MAP_DATA: state => state.MapData,
        COUNTRIES_DATA: state => state.CountriesData,
        TOP_COUNTRIES_DATA: state => state.TopCountriesData,
        SOFT_DATA: state => state.SoftData,
        TOP_PREFIX_DATA: state => state.TopPrefixData,
        PREFIX_DATA: state => state.PrefixData,
        WIN_DATA: state => state.WinData,
    },
    mutations: {
        SET_DAYS_DATA: (state, payload) => {
            state.DaysData = Object.values(payload);
        },
        SET_LOGS_INFO: (state, payload) => {
            state.LogsInfo = payload;
        },
        SET_MAP_DATA: (state, payload) => {
            state.MapData = payload;
        },
        SET_TOP_COUNTRIES_DATA: (state, payload) => {
            state.TopCountriesData = payload;
        },
        SET_COUNTRIES_DATA: (state, payload) => {
            state.CountriesData = payload;
        },
        SET_SOFT_DATA: (state, payload) => {
            state.SoftData = payload;
        },
        SET_TOP_PREFIX_DATA: (state, payload) => {
            state.TopPrefixData = payload;
        },
        SET_PREFIX_DATA: (state, payload) => {
            state.PrefixData = payload;
        },
        SET_WIN_DATA: (state, payload) => {
            state.WinData = payload;
        },
    },
    actions: {
        GET_DAYS_DATA: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/dash/daysData/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_DAYS_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        GET_LOGS_INFO: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/dash/logsInfo/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_LOGS_INFO', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        GET_MAP_DATA: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/dash/mapData/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_MAP_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        GET_TOP_COUNTRIES_DATA: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/dash/topCountriesData/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_TOP_COUNTRIES_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        GET_COUNTRIES_DATA: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/dash/countriesData/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_COUNTRIES_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        GET_SOFT_DATA: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/dash/softData/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_SOFT_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        GET_TOP_PREFIX_DATA: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/dash/topPrefixData/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_TOP_PREFIX_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        GET_PREFIX_DATA: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/dash/prefixData/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_PREFIX_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
        GET_WIN_DATA: async ({commit, rootGetters}) => {
            await client({
                method: 'POST',
                url: '/dash/winData/',
                data: {'cookie': rootGetters.USER_COOKIE},
            }).then(resp => {
                commit('SET_WIN_DATA', resp.data);
            })
            .catch(error => {
                console.error(error);
            });
        },
    }
}