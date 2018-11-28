<template>

<div >
        <div>
          <h3 style="color:black;text-decoration:underline;text-align:center;">Demo App for Message Board</h3>
        </div>

<div class="col-md-10 col-md-offset-1">

                    <vue-good-table
                      :columns="columns"
                      :rows="rows"
                      :pagination-options="pagination_options"
                      :lineNumbers="true" >

                      <template slot="table-column" slot-scope="props">
                        <span class="vue-table-header">
                           {{props.column.label}}
                        </span>
                      </template>

                      <template slot="table-row" slot-scope="props">
                          <span v-if="props.column.field == 'actions'">
                             <button type="button" class="btn btn-sm btn-primary" aria-label="Left Align" @click="itemAction('expand-item',props.row)" ><i class="glyphicon glyphicon-plus-sign " aria-hidden="true"></i></button>
                             <button type="button" class="btn btn-sm btn-danger" aria-label="Left Align" @click="itemAction('delete-item',props.row)" ><i class="glyphicon glyphicon-trash " aria-hidden="true"></i></button>
                          </span>

                          <span v-else class="vue-table-rows">
                            {{props.formattedRow[props.column.field]}}
                          </span>
                      </template>

                      <div slot="table-actions">
                                <button type="button" class="btn btn-primary " @click="showFormModal = true" ><i class="glyphicon glyphicon-send"></i></button>
                                <button type="button" class="btn btn-primary " @click="RefreshData()" ><i class="glyphicon glyphicon-refresh"></i></button>
                      </div>

                  </vue-good-table>

                    <!--Details Modal -->
            <div v-if="showModal">
              <transition name="modal">
                <div class="modal-mask">
                  <div class="modal-wrapper">
                    <div class="modal-dialog">
                      <div class="modal-content">
                        <div class="modal-header">
                          <h3 class="modal-title" >Message Details</h3>
                          <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="showModal =false">
                            <span aria-hidden="true">&times;</span>
                          </button>
                        </div>
                        <div class="modal-body">
                          <div class="col-md-10">
                          <div class="table-responsive">
                          <table class="table table-striped table-hover table-condensed">
                                  <thead>
                                    <tr>
                                      <th>Author</th>
                                      <th>Message Title</th>
                                      <th>Message Description</th>
                                      <th> Pallindrome </th>
                                    </tr>
                                  </thead>
                                  <tbody>
                                    <tr v-for="item in items">
                                      <td>{{item.author}}</td>
                                      <td>{{item.title}}</td>
                                      <td>{{item.description}}</td>
                                      <td>{{item.pallindrome}}</td>
                                    </tr>

                                  </tbody>
                                </table>
                              </div>
                        </div>
                        </div>
                        <div class="modal-footer">
                          <button type="button" class="btn btn-secondary" data-dismiss="modal" @click="showModal =false">Close</button>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </transition>
            </div>


            <!-- Form Modal -->
              <div v-if="showFormModal">
                <transition name="modal">
                  <div class="modal-mask">
                    <div class="modal-wrapper">
                      <div class="modal-dialog">
                        <div class="modal-content">
                          <div class="modal-header">
                            <h3 class="modal-title" >Add Message </h3>
                            <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="showFormModal =false">
                              <span aria-hidden="true">&times;</span>
                            </button>
                          </div>
                          <div class="modal-body">
                            <div class="col-md-10">
                              <form class="row">
                                <div class="form-group col-md-8">
                                        <label for="author" class="form-control-label">Author</label>
                                        <input type="text" id="author" class="form-control" v-model="author" />
                                </div>
                                <div class="form-group col-md-8">
                                        <label for="title" class="form-control-label">Message Title</label>
                                        <input type="text" id="title" class="form-control" v-model="title" />
                                </div>
                                <div class="form-group col-md-8">
                                        <label for="desc`" class="form-control-label">Message Description</label>
                                        <textarea class="form-control" rows="2" v-model="desc"  ></textarea>
                                </div>
                              </form>
                          </div>
                          </div>
                          <div class="modal-footer">
                            <button type="button" class="btn btn-secondary" data-dismiss="modal" @click="showFormModal =false">Close</button>
                            <button type="button" class="btn btn-primary" data-dismiss="modal" @click="addMessage() ">Add Message</button>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </transition>
              </div>

                  <notifications group="app" position="bottom center" width="400"/>
                  <loading :active.sync="loading" color="#00ff80"></loading>

</div>
</div>
</template>




<script>
import { VueGoodTable } from 'vue-good-table'
import Loading from 'vue-loading-overlay'
 import notificate  from '../functions/notifications'
 import {apiService} from '../functions/apiServices'

 export default {
data:function (){
    return{
      url :"",
      pagination_options:{
        enabled:true,
        perPage:8,
        mode: 'pages'
      },
      columns: [
        {
          label:'id',
          field:'id',
          hidden:'true'
        },
        {
          label:'Author',
          field:'author',
        },
        {
          label:'Message Title',
          field:'title',
        },
        {
          label:'Actions',
          field:'actions'
        },
      ],
      rows:[],
      loading:false,
      showModal:false,
      showFormModal:false,
      title:'',
      author:'',
      desc:'',
      item:[],
    }
  },
  methods:{
    async RefreshData(){
      var apiUrl =  this.url+"/messages"
      this.loading = true
      var status =await apiService.Get(apiUrl)
      this.loading = false
      if(status.error){
          return this.notificate('App says',status.error,'error')
      }
      if (status.data == null){
        return this.notificate('App says','No data in table right now','warning')
      }
      this.rows = status.data

    },

    async deleteMessage(id){
      var apiUrl = this.url+"/message/"+id
      this.loading=true
      var status = await apiService.Delete(apiUrl);
      this.loading=false
      if(status.error){
            return this.notificate('App says',status.error,'error')
        }
        this.notificate('App says','Message Deleted','success')
    },

    async addMessage(){
      if(this.title ==""||this.author ==""||this.desc==""){
        return this.notificate('app says',"All Fields Are required",'error')
      }
        var data={
        'author':this.author,
        'title' : this.title,
        'body': this.desc
      }

      var apiUrl = this.url+"/message"
      this.showFormModal=false
      this.loading=true
      var status = await apiService.Post(apiUrl,JSON.stringify(data));
      this.loading=false
      if(status.error){
            return this.notificate('App says',status.error,'error')
        }
        await this.RefreshData()
        this.notificate('App says','Message Added','success')
    },

    async getDetail(id){
        var apiUrl = this.url+"/message/"+id
        this.loading = true
        var status =await apiService.Get(apiUrl)
        this.loading = false
        if(status.error){
            return this.notificate('App says',status.error,'error')
        }
        this.items = status.data
        this.showModal=true
    },
    async itemAction(action,data){
      if(action =="expand-item"){
          await this.getDetail(data.id)
      }else if (action =="delete-item") {
          await this.deleteMessage(data.id)
          await this.RefreshData()
      }
    }
  },
watch:{
      showFormModal:function(){
        if(this.showFormModal == false){
          this.title=''
          this.author=''
          this.desc =''
        }
      }
  },
  created(){
     this.RefreshData()
  },
components:{
    VueGoodTable,
    Loading,

  },
mixins:[notificate]
 }
</script>
