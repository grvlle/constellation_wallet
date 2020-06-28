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
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}