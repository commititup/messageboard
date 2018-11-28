export  default {
 methods: {
   notificate (title,text,type){
        this.$notify({
                group: 'app',
                title: title,
                text: text,
                type: type,
                speed: 300,
                duration:3000
              })
    }
  }
}
