/*
对 storage 进行二次封装
*/
import config from "./../config"

export default {
    getStorage(){
        return JSON.parse(window.localStorage.getItem(config.namespace) || '{}')
    },

    setItem(key, value){
        let storage = this.getStorage()
        storage[key] = value
        window.localStorage.setItem(config.namespace, JSON.stringify(storage))
    },
    getItem(key){
        return this.getStorage()[key]
    },
    clearItem(key){
        let storage = this.getStorage()
        delete storage[key]
        window.localStorage.setItem(config.namespace, JSON.stringify(storage))
    },  
    clearAll(){
        window.localStorage.clear()
    }
}