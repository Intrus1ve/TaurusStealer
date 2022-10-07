import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex);

import login from './login'
import dash from './dash'
import logs from './logs'
import grabber from './grabber'
import loader from './loader'
import config from './config'
import users from './users'
import backups from './backups'
import dd from './dd'
import banned from './banned'

let store = new Vuex.Store({
    state: {
        LastError: null,
    },
    getters: {
        GET_LAST_ERROR: state => {
            return state.LastError;
        },
    },
    mutations: {
        SET_LAST_ERROR: (state, payload) => {
            state.LastError = payload;
        },
    },
    actions: {
    },
    modules: {
        login, dash, logs, grabber, loader, config, users, backups, dd, banned
    }
});

export default store;