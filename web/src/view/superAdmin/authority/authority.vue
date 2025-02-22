<template>
  <div class="authority">
    <warning-bar title="Note: The upper right corner is pulled down to switch the character" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="addAuthority(0)">New Role</el-button>
      </div>
      <el-table :data="tableData" :tree-props="{ children: 'children', hasChildren: 'hasChildren' }" row-key="authorityId"
        style="width: 100%">
        <el-table-column label="ID" min-width="180" prop="authorityId" />
        <el-table-column align="left" label="Role Name" min-width="180" prop="authorityName" />
        <el-table-column align="left" label="Action" width="460">
          <template #default="scope">
            <el-button icon="setting" size="small" type="primary" link @click="opdendrawer(scope.row)">Setting
              Permission</el-button>
            <el-button icon="plus" size="small" type="primary" link @click="addAuthority(scope.row.authorityId)">Add Authority
            </el-button>
            <el-button icon="copy-document" size="small" type="primary" link @click="copyAuthorityFunc(scope.row)">Copy
            </el-button>
            <el-button icon="edit" size="small" type="primary" link @click="editAuthority(scope.row)">Edit</el-button>
            <el-button icon="delete" size="small" type="primary" link @click="deleteAuth(scope.row)">Delete</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- Add new角色弹窗 -->
    <el-dialog v-model="dialogFormVisible" :title="dialogTitle">
      <el-form ref="authorityForm" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="Parent" prop="parentId">
          <el-cascader v-model="form.parentId" style="width:100%" :disabled="dialogType == 'add'"
            :options="AuthorityOption"
            :props="{ checkStrictly: true, label: 'authorityName', value: 'authorityId', disabled: 'disabled', emitPath: false }"
            :show-all-levels="false" filterable />
        </el-form-item>
        <el-form-item label="ID" prop="authorityId">
          <el-input v-model="form.authorityId" :disabled="dialogType == 'edit'" autocomplete="off" />
        </el-form-item>
        <el-form-item label="Name" prop="authorityName">
          <el-input v-model="form.authorityName" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">Cancel</el-button>
          <el-button size="small" type="primary" @click="enterDialog">Sure</el-button>
        </div>
      </template>
    </el-dialog>

    <el-drawer v-if="drawer" v-model="drawer" custom-class="auth-drawer" :with-header="false" size="40%" title="configuration">
      <el-tabs :before-leave="autoEnter" type="border-card">
        <el-tab-pane label="Menu">
          <Menus ref="menus" :row="activeRow" @changeRow="changeRow" />
        </el-tab-pane>
        <el-tab-pane label="API">
          <Apis ref="apis" :row="activeRow" @changeRow="changeRow" />
        </el-tab-pane>
        <el-tab-pane label="Resource Permission">
          <Datas ref="datas" :authority="tableData" :row="activeRow" @changeRow="changeRow" />
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  getAuthorityList,
  deleteAuthority,
  createAuthority,
  updateAuthority,
  copyAuthority
} from '@/api/authority'

import Menus from '@/view/superAdmin/authority/components/menus.vue'
import Apis from '@/view/superAdmin/authority/components/apis.vue'
import Datas from '@/view/superAdmin/authority/components/datas.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'

import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const mustUint = (rule, value, callback) => {
  if (!/^[0-9]*[1-9][0-9]*$/.test(value)) {
    return callback(new Error('Please enter the positive integer'))
  }
  return callback()
}

const AuthorityOption = ref([
  {
    authorityId: 0,
    authorityName: 'Root'
  }
])
const drawer = ref(false)
const dialogType = ref('add')
const activeRow = ref({})

const dialogTitle = ref('New role')
const dialogFormVisible = ref(false)
const apiDialogFlag = ref(false)
const copyForm = ref({})

const form = ref({
  authorityId: 0,
  authorityName: '',
  parentId: 0
})
const rules = ref({
  authorityId: [
    { required: true, message: 'Please enter the character ID', trigger: 'blur' },
    { validator: mustUint, trigger: 'blur', message: 'Must be a positive integer' }
  ],
  authorityName: [
    { required: true, message: 'Please enter the character name', trigger: 'blur' }
  ],
  parentId: [
    { required: true, message: 'Please choose a parent role', trigger: 'blur' },
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(999)
const tableData = ref([])
const searchInfo = ref({})

// 查询
const getTableData = async () => {
  const table = await getAuthorityList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const changeRow = (key, value) => {
  activeRow.value[key] = value
}
const menus = ref(null)
const apis = ref(null)
const datas = ref(null)
const autoEnter = (activeName, oldActiveName) => {
  const paneArr = [menus, apis, datas]
  if (oldActiveName) {
    if (paneArr[oldActiveName].value.needConfirm) {
      paneArr[oldActiveName].value.enterAndNext()
      paneArr[oldActiveName].value.needConfirm = false
    }
  }
}
// 拷贝角色
const copyAuthorityFunc = (row) => {
  setOptions()
  dialogTitle.value = 'Copy'
  dialogType.value = 'copy'
  for (const k in form.value) {
    form.value[k] = row[k]
  }
  copyForm.value = row
  dialogFormVisible.value = true
}
const opdendrawer = (row) => {
  drawer.value = true
  activeRow.value = row
}
// Delete角色
const deleteAuth = (row) => {
  ElMessageBox.confirm('Please choose a parent role', 'Hint', {
    confirmButtonText: 'Sure',
    cancelButtonText: 'Cancel',
    type: 'warning'
  })
    .then(async () => {
      const res = await deleteAuthority({ authorityId: row.authorityId })
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
    .catch(() => {
      ElMessage({
        type: 'info',
        message: 'Cancelled'
      })
    })
}
// 初始化表单
const authorityForm = ref(null)
const initForm = () => {
  if (authorityForm.value) {
    authorityForm.value.resetFields()
  }
  form.value = {
    authorityId: 0,
    authorityName: '',
    parentId: 0
  }
}
// 关闭窗口
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
  apiDialogFlag.value = false
}
// Sure弹窗

const enterDialog = () => {
  form.value.authorityId = Number(form.value.authorityId)
  if (form.value.authorityId === 0) {
    ElMessage({
      type: 'error',
      message: 'The character ID cannot be 0'
    })
    return false
  }
  authorityForm.value.validate(async valid => {
    if (valid) {
      switch (dialogType.value) {
        case 'add':
          {
            const res = await createAuthority(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: 'Added successfully!'
              })
              getTableData()
              closeDialog()
            }
          }
          break
        case 'edit':
          {
            const res = await updateAuthority(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: 'Added successfully!'
              })
              getTableData()
              closeDialog()
            }
          }
          break
        case 'copy': {
          const data = {
            authority: {
              authorityId: 0,
              authorityName: '',
              datauthorityId: [],
              parentId: 0
            },
            oldAuthorityId: 0
          }
          data.authority.authorityId = form.value.authorityId
          data.authority.authorityName = form.value.authorityName
          data.authority.parentId = form.value.parentId
          data.authority.dataAuthorityId = copyForm.value.dataAuthorityId
          data.oldAuthorityId = copyForm.value.authorityId
          const res = await copyAuthority(data)
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: 'Copy successfully!'
            })
            getTableData()
          }
        }
      }

      initForm()
      dialogFormVisible.value = false
    }
  })
}
const setOptions = () => {
  AuthorityOption.value = [
    {
      authorityId: 0,
      authorityName: 'Root'
    }
  ]
  setAuthorityOptions(tableData.value, AuthorityOption.value, false)
}
const setAuthorityOptions = (AuthorityData, optionsData, disabled) => {
  form.value.authorityId = String(form.value.authorityId)
  AuthorityData &&
    AuthorityData.forEach(item => {
      if (item.children && item.children.length) {
        const option = {
          authorityId: item.authorityId,
          authorityName: item.authorityName,
          disabled: disabled || item.authorityId === form.value.authorityId,
          children: []
        }
        setAuthorityOptions(
          item.children,
          option.children,
          disabled || item.authorityId === form.value.authorityId
        )
        optionsData.push(option)
      } else {
        const option = {
          authorityId: item.authorityId,
          authorityName: item.authorityName,
          disabled: disabled || item.authorityId === form.value.authorityId
        }
        optionsData.push(option)
      }
    })
}
// 增加角色
const addAuthority = (parentId) => {
  initForm()
  dialogTitle.value = 'New Role'
  dialogType.value = 'add'
  form.value.parentId = parentId
  setOptions()
  dialogFormVisible.value = true
}
// 编辑角色
const editAuthority = (row) => {
  setOptions()
  dialogTitle.value = 'Edit'
  dialogType.value = 'edit'
  for (const key in form.value) {
    form.value[key] = row[key]
  }
  setOptions()
  dialogFormVisible.value = true
}

</script>

<script>

export default {
  name: 'Authority'
}
</script>

<style lang="scss">
.authority {
  .el-input-number {
    margin-left: 15px;

    span {
      display: none;
    }
  }
}

.tree-content {
  overflow: auto;
  height: calc(100vh - 100px);
  margin-top: 10px;
}

.auth-drawer {
  .el-drawer__body {
    overflow: hidden;
  }
}
</style>
