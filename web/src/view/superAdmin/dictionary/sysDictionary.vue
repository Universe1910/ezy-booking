<template>
  <div>
    <warning-bar
      title="Get the dictionary and the cache method has been encapsulated at the front Utils/Dictionary."
    />
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="Dictionary name (middle)">
          <el-input v-model="searchInfo.name" placeholder="Search condition" />
        </el-form-item>
        <el-form-item label="Dictionary name (English)">
          <el-input v-model="searchInfo.type" placeholder="Search condition" />
        </el-form-item>
        <el-form-item label="State" prop="status">
          <el-select v-model="searchInfo.status" clear placeholder="Please choose">
            <el-option key="true" label="Yes" value="true" />
            <el-option key="false" label="No" value="false" />
          </el-select>
        </el-form-item>
        <el-form-item label="Description">
          <el-input v-model="searchInfo.desc" placeholder="Search condition" />
        </el-form-item>
        <el-form-item>
          <el-button
            size="small"
            type="primary"
            icon="search"
            @click="onSubmit"
          >Search</el-button>
          <el-button
            size="small"
            icon="refresh"
            @click="onReset"
          >Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          size="small"
          type="primary"
          icon="plus"
          @click="openDialog"
        >Add New</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="Date" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>

        <el-table-column
          align="left"
          label="Name (middle)"
          prop="name"
          width="160"
        />

        <el-table-column
          align="left"
          label="Name (English)"
          prop="type"
          width="120"
        />

        <el-table-column align="left" label="State" prop="status" width="120">
          <template #default="scope">{{
            formatBoolean(scope.row.status)
          }}</template>
        </el-table-column>

        <el-table-column align="left" label="Describe" prop="desc" width="280" />

        <el-table-column align="left" label="Action">
          <template #default="scope">
            <el-button
              size="small"
              icon="document"
              type="primary"
              link
              @click="toDetail(scope.row)"
            >Detail</el-button>
            <el-button
              size="small"
              icon="edit"
              type="primary"
              link
              @click="updateSysDictionaryFunc(scope.row)"
            >Change</el-button>
            <el-popover
              v-model="scope.row.visible"
              placement="top"
              width="160"
            >
              <p>You sure you want to delete it?？</p>
              <div style="text-align: right; margin-top: 8px">
                <el-button
                  size="small"
                  type="primary"
                  link
                  @click="scope.row.visible = false"
                >Cancel</el-button>
                <el-button
                  type="primary"
                  size="small"
                  @click="deleteSysDictionaryFunc(scope.row)"
                >Sure</el-button>
              </div>
              <template #reference>
                <el-button
                  type="primary"
                  link
                  icon="delete"
                  size="small"
                  style="margin-left: 10px"
                  @click="scope.row.visible = true"
                >Delete</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>

      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      title="Pop-up operation"
    >
      <el-form
        ref="dialogForm"
        :model="formData"
        :rules="rules"
        size="medium"
        label-width="110px"
      >
        <el-form-item label="Name (middle)" prop="name">
          <el-input
            v-model="formData.name"
            placeholder="Please enter the dictionary name (middle)"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item label="Name (English)" prop="type">
          <el-input
            v-model="formData.type"
            placeholder="Please enter the dictionary name (English)"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item label="State" prop="status" required>
          <el-switch
            v-model="formData.status"
            active-text="Open"
            inactive-text="Close"
          />
        </el-form-item>
        <el-form-item label="describe" prop="desc">
          <el-input
            v-model="formData.desc"
            placeholder="Describe"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">Cancel</el-button>
          <el-button
            size="small"
            type="primary"
            @click="enterDialog"
          >Sure</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'SysDictionary',
}
</script>

<script setup>
import {
  createSysDictionary,
  deleteSysDictionary,
  updateSysDictionary,
  findSysDictionary,
  getSysDictionaryList,
} from '@/api/sysDictionary' // 此处请自行替换地址
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { formatBoolean, formatDate } from '@/utils/format'

const router = useRouter()

const formData = ref({
  name: null,
  type: null,
  status: true,
  desc: null,
})

const rules = ref({
  name: [
    {
      required: true,
      message: 'Please enter the dictionary name (middle)',
      trigger: 'blur',
    },
  ],
  type: [
    {
      required: true,
      message: 'Please enter the dictionary name (English)',
      trigger: 'blur',
    },
  ],
  desc: [
    {
      required: true,
      message: 'Please enter describe',
      trigger: 'blur',
    },
  ],
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
}

// 条件搜索前端看此方法
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.status === '') {
    searchInfo.value.status = null
  }
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

// 查询
const getTableData = async() => {
  const table = await getSysDictionaryList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const toDetail = (row) => {
  router.push({
    name: 'dictionaryDetail',
    params: {
      id: row.ID,
    },
  })
}

const dialogFormVisible = ref(false)
const type = ref('')
const updateSysDictionaryFunc = async(row) => {
  const res = await findSysDictionary({ ID: row.ID, status: row.status })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.resysDictionary
    dialogFormVisible.value = true
  }
}
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: null,
    type: null,
    status: true,
    desc: null,
  }
}
const deleteSysDictionaryFunc = async(row) => {
  row.visible = false
  const res = await deleteSysDictionary({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Delete成功',
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

const dialogForm = ref(null)
const enterDialog = async() => {
  dialogForm.value.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createSysDictionary(formData.value)
        break
      case 'update':
        res = await updateSysDictionary(formData.value)
        break
      default:
        res = await createSysDictionary(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage.success('Successful operation')
      closeDialog()
      getTableData()
    }
  })
}
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}
</script>

<style></style>
