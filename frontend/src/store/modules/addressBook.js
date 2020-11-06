const getDefaultState = () => {
    return {
        addressBook: []
    }
}

const state = getDefaultState()

const getters = {
    search: (state) => (searchText) => {
        return state.addressBook.filter(function (contact) {
            if (contact.name.toUpperCase().includes(searchText.toUpperCase()) ||
                contact.address.includes(searchText) ||
                contact.tag.toUpperCase().includes(searchText.toUpperCase())
            ) {
                return true;
            } else {
                return false;
            }
        });
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
        const index = state.addressBook.findIndex((contact) => contact.id === obj.contact.id);
        if (index === -1) {
            state.addressBook.push(obj.contact);
        } else {
            state.addressBook[index] = obj.contact;
        }
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