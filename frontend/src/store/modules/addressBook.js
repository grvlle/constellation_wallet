const getDefaultState = () => {
    return {
        addressBook: []
    }
}

const state = getDefaultState()

const getters = {
    search: (state) => (searchText) => {
        return state.addressBook.filter(contact => {
            return contact.name.toUpperCase().includes(searchText.toUpperCase())
        })
    },
    byId: (state) => (id) => {
        return state.addressBook.find(contact => contact.id === id)
    }
}

const actions = {
    reset({ commit }) {
        commit('resetState')
    }
}
const mutations = {
    resetState(state) {
        Object.assign(state, getDefaultState())
    },
    setAddressBook(state, obj) {
        state.addressBook = obj.addressBook;
    },
    setContact(state, obj) {
        state.addressBook.push(obj.contact);
    },
    deleteContact(state, obj) {
        state.addressBook = state.addressBook.filter(contact => contact.id === obj.contact.id);
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}