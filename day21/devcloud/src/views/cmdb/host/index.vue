<template>
  <div class="host-container">
    <tips :tips="tips" />
    <div class="table-op">
      <div class="search">
        <el-input placeholder="请输入内容" v-model="filterValue" class="input-with-select">
          <el-select v-model="filterKey" slot="prepend" placeholder="请选择">
            <el-option label="餐厅名" value="1"></el-option>
            <el-option label="订单号" value="2"></el-option>
            <el-option label="用户电话" value="3"></el-option>
          </el-select>
          <el-button slot="append" icon="el-icon-search"></el-button>
        </el-input>
      </div>
      <div class="op">

      </div>
    </div>
    <div class="box-shadow">
      <el-table
        :data="hosts"
        style="width: 100%">
        <el-table-column
          prop="name"
          label="实例名称">
          <template slot-scope="scope">
            {{ scope.row.instance_id}}
            <br>
            {{ scope.row.name}} 
          </template>
        </el-table-column>
        <el-table-column
          prop="name"
          label="资产来源">
          <template slot-scope="scope">
            {{ scope.row.vendor}}
            <br>
            {{ scope.row.region}} 
          </template>
        </el-table-column>
        <el-table-column
          prop="private_ip"
          label="内网IP">
          <template slot-scope="scope">
            {{ scope.row.private_ip}} 
          </template>
        </el-table-column>
        <el-table-column
          prop="os_name"
          label="系统类型">
          <template slot-scope="scope">
            {{ scope.row.os_name}} 
          </template>
        </el-table-column>
        <el-table-column
          prop="create_at"
          label="创建时间">
          <template slot-scope="scope">
            {{ scope.row.create_at | parseTime}}
          </template>
        </el-table-column>
        <el-table-column
          prop="expire_at"
          label="过期时间">
          <template slot-scope="scope">
            {{ scope.row.expire_at | parseTime}}
          </template>
        </el-table-column>
        <el-table-column
          prop="create_at"
          label="配置规格">
          <template slot-scope="scope">
            {{ scope.row.cpu}} / {{ scope.row.memory }}
          </template>
        </el-table-column>
        <el-table-column
          prop="expire_at"
          label="状态">
          <template slot-scope="scope">
            {{ scope.row.status}}
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          align="center"
          width="180">
          <template slot-scope="scope">
            <el-button type="text" disabled>归档</el-button>
            <el-button type="text" @click="jumpToMonitor(scope.row.instance_id)">监控</el-button>
          </template>
        </el-table-column>
      </el-table>

      <pagination 
        v-show="total>0" 
        :total="total" 
        :page.sync="query.page_number" 
        :limit.sync="query.page_size" 
        @pagination="get_hosts" 
      />
    </div>
  </div>
</template>

<script>
import Tips from '@/components/Tips'
import Pagination from '@/components/Pagination'
import { LIST_HOST } from '@/api/cmdb/host.js'

const tips = [
  '现在仅同步了阿里云主机资产'
]

export default {
  name: 'CmdbHost',
  components: { Tips, Pagination },
  data() {
    return {
      tips: tips,
      filterKey: '',
      filterValue: '',
      query: {page_size: 20, page_number: 1},
      total: 0,
      hosts: []
    }
  },
  created() {
    this.get_hosts()
  },
  methods: {
    async get_hosts() {
      const resp = await LIST_HOST(this.query)
      console.log(resp)
      this.hosts = resp.data.items
      this.total = resp.data.total
    },
    handleSizeChange(val) {
      this.query.page_size = val
      this.get_hosts()
    },
    handleCurrentChange(val) {
      this.query.page_number = val
      this.get_hosts()
    },
    jumpToMonitor(id) {
      console.log(id)
    }
  }
}
</script>

<style lang="scss" scoped>
.box-shadow {
  margin: 12px 0;
}
</style>