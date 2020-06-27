const getDefaultState = () => {
    return {
        addressBook: [
            { id: 1, name: "John Doe", tag: "PERSONAL" },
            { id: 2, name: "Jane Doe", tag: "PERSONAL" }
        ],
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