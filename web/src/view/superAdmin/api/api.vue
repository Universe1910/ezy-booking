<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="Path">
          <el-input v-model="searchInfo.path" placeholder="Path" />
        </el-form-item>
        <el-form-item label="Describe">
          <el-input v-model="searchInfo.description" placeholder="describe" />
        </el-form-item>
        <el-form-item label="API Group">
          <el-input v-model="searchInfo.apiGroup" placeholder="api Group" />
        </el-form-item>
        <el-form-item label="Method">
          <el-select v-model="searchInfo.method" clearable placeholder="Choose">
            <el-option v-for="item in methodOptions" :key="item.value" :label="`${item.label}(${item.value})`"
              :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">Search</el-button>
          <el-button size="small" icon="refresh" @click="onReset">Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog('addApi')">Add new</el-button>
        <el-popover v-model="deleteVisible" placement="top" width="160">
          <p>You sure you want to delete it?</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">Cancel</el-button>
            <el-button size="small" type="primary" @click="onDelete">Sure</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" :disabled="!apis.length" style="margin-left: 10px;"
              @click="deleteVisible = true">Delete</el-button>
          </template>
        </el-popover>
      </div>
      <el-table :data="tableData" @sort-change="sortChange" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="id" min-width="60" prop="ID" sortable="custom" />
        <el-table-column align="left" label="API Path" min-width="150" prop="path" sortable="custom" />
        <el-table-column align="left" label="API Group" min-width="150" prop="apiGroup" sortable="custom" />
        <el-table-column align="left" label="API Description" min-width="150" prop="description" sortable="custom" />
        <el-table-column align="left" label="Method" min-width="150" prop="method" sortable="custom">
          <template #default="scope">
            <div>
              {{ scope.row.method }} / {{ methodFilter(scope.row.method) }}
            </div>
          </template>
        </el-table-column>

        <el-table-column align="left" fixed="right" label="Action" width="200">
          <template #default="scope">
            <el-button icon="edit" size="small" type="primary" link @click="editApiFunc(scope.row)">Edit</el-button>
            <el-button icon="delete" size="small" type="primary" link @click="deleteApiFunc(scope.row)">Delete
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
          layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>

    </div>

    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle">
      <warning-bar title="Add newAPI，you need to configure the permissions in the character management to use" />
      <el-form ref="apiForm" :model="form" :rules="rules" label-width="120px">
        <el-form-item label="Path" prop="path">
          <el-input v-model="form.path" autocomplete="off" />
        </el-form-item>
        <el-form-item label="Method" prop="method">
          <el-select v-model="form.method" placeholder="Choose" style="width:100%">
            <el-option v-for="item in methodOptions" :key="item.value" :label="`${item.label}(${item.value})`"
              :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="api group" prop="apiGroup">
          <el-input v-model="form.apiGroup" autocomplete="off" />
        </el-form-item>
        <el-form-item label="api description" prop="description">
          <el-input v-model="form.description" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">Cancel</el-button>
          <el-button size="small" type="primary" @click="enterDialog">Sure</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Api',
}
</script>

<script setup>
import {
  getApiById,
  getApiList,
  createApi,
  updateApi,
  deleteApi,
  deleteApisByIds
} from '@/api/api'
import { toSQLLine } from '@/utils/stringFun'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const methodFilter = (value) => {
  const target = methodOptions.value.filter(item => item.value === value)[0]
  return target && `${target.label}`
}

const apis = ref([])
const form = ref({
  path: '',
  apiGroup: '',
  method: '',
  description: ''
})
const methodOptions = ref([
  {
    value: 'POST',
    label: 'Create',
    type: 'success'
  },
  {
    value: 'GET',
    label: 'Check',
    type: ''
  },
  {
    value: 'PUT',
    label: 'Renew',
    type: 'warning'
  },
  {
    value: 'DELETE',
    label: 'Delete',
    type: 'danger'
  }
])

const type = ref('')
const rules = ref({
  path: [{ required: true, message: 'Please enter the API path', trigger: 'blur' }],
  apiGroup: [
    { required: true, message: 'Please enter the name of the group', trigger: 'blur' }
  ],
  method: [
    { required: true, message: 'Please select the request method', trigger: 'blur' }
  ],
  description: [
    { required: true, message: 'Please enter the API introduction', trigger: 'blur' }
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
}
// 搜索

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 排序
const sortChange = ({ prop, order }) => {
  if (prop) {
    if (prop === 'ID') {
      prop = 'id'
    }
    searchInfo.value.orderKey = toSQLLine(prop)
    searchInfo.value.desc = order === 'descending'
  }
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getApiList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// 批量Action
const handleSelectionChange = (val) => {
  apis.value = val
}

const deleteVisible = ref(false)
const onDelete = async () => {
  const ids = apis.value.map(item => item.ID)
  const res = await deleteApisByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: res.msg
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 弹窗相关
const apiForm = ref(null)
const initForm = () => {
  apiForm.value.resetFields()
  form.value = {
    path: '',
    apiGroup: '',
    method: '',
    description: ''
  }
}

const dialogTitle = ref('Add newApi')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
  switch (key) {
    case 'addApi':
      dialogTitle.value = 'Add new API'
      break
    case 'edit':
      dialogTitle.value = 'Edit Api'
      break
    default:
      break
  }
  type.value = key
  dialogFormVisible.value = true
}
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}

const editApiFunc = async (row) => {
  const res = await getApiById({ id: row.ID })
  form.value = res.data.api
  openDialog('edit')
}

const enterDialog = async () => {
  apiForm.value.validate(async valid => {
    if (valid) {
      switch (type.value) {
        case 'addApi':
          {
            const res = await createApi(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: 'Added successfully',
                showClose: true
              })
            }
            getTableData()
            closeDialog()
          }

          break
        case 'edit':
          {
            const res = await updateApi(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: 'Edit',
                showClose: true
              })
            }
            getTableData()
            closeDialog()
          }
          break
        default:
          // eslint-disable-next-line no-lone-blocks
          {
            ElMessage({
              type: 'error',
              message: 'Unknown operation',
              showClose: true
            })
          }
          break
      }
    }
  })
}

const deleteApiFunc = async (row) => {
  ElMessageBox.confirm('This operation will continue the API under all characters of permanent Delete?', 'hint', {
    confirmButtonText: 'Sure',
    cancelButtonText: 'Cancel',
    type: 'warning'
  })
    .then(async () => {
      const res = await deleteApi(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: 'Deleted successfully!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
}

</script>

<style scoped lang="scss">
.button-box {
  padding: 10px 20px;

  .el-button {
    float: right;
  }
}

.warning {
  color: #dc143c;
}
</style>
