<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <!-- <el-form-item label="Create 时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item> -->
        <el-form-item label="Name">
          <el-input v-model="searchInfo.name" placeholder="Stage" />

        </el-form-item>
        <!-- <el-form-item label="Address">
          <el-input v-model="searchInfo.address" placeholder="搜索条件" />
        </el-form-item> -->
        <!-- <el-form-item label="Map URL">
          <el-input v-model="searchInfo.mapURL" placeholder="搜索条件" />
        </el-form-item> -->
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">Search</el-button>
          <el-button size="small" icon="refresh" @click="onReset">Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">Add new</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>Are you want to delete it?</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">Cancel</el-button>
            <el-button size="small" type="primary" @click="onDelete">Delete</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length"
              @click="deleteVisible = true">Delete </el-button>
          </template>
        </el-popover>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="Created At" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Name" prop="name" width="300" />
        <el-table-column align="left" label="Address" prop="address" width="200" />
        <!-- <el-table-column align="left" label="Map URL" prop="mapURL" width="120" /> -->
        <el-table-column align="left" label="Action">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button"
              @click="updateEzyStageFunc(scope.row)">Edit</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">Delete </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="Add new">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="Name:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="Please enter stage name" />
        </el-form-item>
        <el-form-item label="Address:" prop="address">
          <el-input v-model="formData.address" :clearable="false" placeholder="Please enter address" />
        </el-form-item>
        <el-form-item label="Map URL:" prop="mapURL">
          <el-input v-model="formData.mapURL" :clearable="true" placeholder="Please enter map URL" />
        </el-form-item>
        <el-form-item label="Branch:" prop="branchID">
          <!-- <el-input v-model="formData.branchID" :clearable="true" placeholder="Please enter map URL" /> -->
          <el-select v-model="formData.branchID" clearable placeholder="Select">
            <el-option v-for="item in branchs" :key="item.ID" :label="item.name" :value="item.ID" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">Cancel</el-button>
          <el-button size="small" type="primary" @click="enterDialog">Save</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'EzyStage'
}
</script>

<script setup>
import {
  createEzyStage,
  deleteEzyStage,
  deleteEzyStageByIds,
  updateEzyStage,
  findEzyStage,
  getEzyStageList
} from '@/api/ezyStage'

import {
  getEzyBranchList
} from '@/api/ezyBranch'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: '',
  address: '',
  mapURL: '',
  branchID: ''
})

// 验证规则
const rule = reactive({
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  address: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const selectBranch = ref()
// 重置
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

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getEzyStageList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    console.log(tableData.value)
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()


const branchs = ref([])

const getBranchs = async () => {
  const table = await getEzyBranchList()
  if (table.code === 0) {
    branchs.value = table.data.list
    // console.log(branchs.value)
  }
}

getBranchs()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// Delete 行
const deleteRow = (row) => {
  ElMessageBox.confirm('Are you sure you want to delete?', 'hint', {
    confirmButtonText: 'Delete',
    cancelButtonText: 'Cancel',
    type: 'warning'
  }).then(() => {
    deleteEzyStageFunc(row)
  })
}


// 批量Delete 控制标记
const deleteVisible = ref(false)

// 多选Delete 
const onDelete = async () => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: 'Please select the data to delete'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.ID)
    })
  const res = await deleteEzyStageByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'successfully deleted'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateEzyStageFunc = async (row) => {
  const res = await findEzyStage({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reezyStage
    dialogFormVisible.value = true
  }
}


// Delete 行
const deleteEzyStageFunc = async (row) => {
  const res = await deleteEzyStage({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'successfully deleted'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: '',
    address: '',
    mapURL: '',
    branchID: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    debugger;
    // formData.value.branchID = selectBranch.value
    switch (type.value) {
      case 'create':
        
        res = await createEzyStage(formData.value)
        break
      case 'update':
        res = await updateEzyStage(formData.value)
        break
      default:
        res = await createEzyStage(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: 'Create sucessfully'
      })
      closeDialog()
      getTableData()
    }
  })
}
</script>

<style>

</style>
