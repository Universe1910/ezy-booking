<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="Create 时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要Delete 吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button size="small" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">Delete </el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="Qty" prop="quantity" width="120" />
        <el-table-column align="left" label="Seats" prop="seats" width="120" />
        <el-table-column align="left" label="Total" prop="total" width="120" />
        <el-table-column align="left" label="Coupon ID" prop="couponID" width="120" />
        <el-table-column align="left" label="Discount" prop="couponDiscount" width="120" />
        <el-table-column align="left" label="Tax" prop="tax" width="120" />
        <el-table-column align="left" label="Invoice Payment" prop="invoicePayment" width="120" />
        <el-table-column align="left" label="Payment Type" prop="paymentType" width="120" />
        <el-table-column align="left" label="Admin Discount" prop="adminDiscount" width="120" />
        <el-table-column align="left" label="Arrived Number" prop="arrivedNumber" width="120" />
        <el-table-column align="left" label="BusNumber" prop="busNumber" width="120" />
        <el-table-column align="left" label="Created By" prop="createdBy" width="120" />
        <el-table-column align="left" label="Last Edit By" prop="lastEditBy" width="120" />
         <el-table-column align="left" label="Last Action Time" width="180">
            <template #default="scope">{{ formatDate(scope.row.lastActionTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="Source" prop="source" width="120" />
        <el-table-column align="left" label="Notes" prop="notes" width="120" />
        <el-table-column align="left" label="Status" prop="status" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateEzyOrdersFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">Delete </el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="Qty:"  prop="quantity" >
          <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Seats:"  prop="seats" >
          <el-input v-model="formData.seats" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Total:"  prop="total" >
          <el-input-number v-model="formData.total"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="Coupon ID:"  prop="couponID" >
          <el-input v-model.number="formData.couponID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Discount:"  prop="couponDiscount" >
          <el-input-number v-model="formData.couponDiscount"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="Tax:"  prop="tax" >
          <el-input-number v-model="formData.tax"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="Invoice Payment:"  prop="invoicePayment" >
          <el-input-number v-model="formData.invoicePayment"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="Payment Type:"  prop="paymentType" >
          <el-input v-model="formData.paymentType" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Admin Discount:"  prop="adminDiscount" >
          <el-input-number v-model="formData.adminDiscount"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="Arrived Number:"  prop="arrivedNumber" >
          <el-input v-model.number="formData.arrivedNumber" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="BusNumber:"  prop="busNumber" >
          <el-input v-model.number="formData.busNumber" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Created By:"  prop="createdBy" >
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Last Edit By:"  prop="lastEditBy" >
          <el-input v-model.number="formData.lastEditBy" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Last Action Time:"  prop="lastActionTime" >
          <el-date-picker v-model="formData.lastActionTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="Source:"  prop="source" >
          <el-input v-model="formData.source" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Notes:"  prop="notes" >
          <el-input v-model="formData.notes" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Status:"  prop="status" >
          <el-input v-model="formData.status" :clearable="false"  placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'EzyOrders'
}
</script>

<script setup>
import {
  createEzyOrders,
  deleteEzyOrders,
  deleteEzyOrdersByIds,
  updateEzyOrders,
  findEzyOrders,
  getEzyOrdersList
} from '@/api/ezyOrders'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        quantity: 0,
        seats: '',
        total: 0,
        couponID: 0,
        couponDiscount: 0,
        tax: 0,
        invoicePayment: 0,
        paymentType: '',
        adminDiscount: 0,
        arrivedNumber: 0,
        busNumber: 0,
        createdBy: 0,
        lastEditBy: 0,
        lastActionTime: new Date(),
        source: '',
        notes: '',
        status: '',
        })

// 验证规则
const rule = reactive({
               status : [{
                   required: true,
                   message: 'Status can not be empty',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

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
const getTableData = async() => {
  const table = await getEzyOrdersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
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
    ElMessageBox.confirm('确定要Delete 吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteEzyOrdersFunc(row)
        })
    }


// 批量Delete 控制标记
const deleteVisible = ref(false)

// 多选Delete 
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要Delete 的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteEzyOrdersByIds({ ids })
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
const updateEzyOrdersFunc = async(row) => {
    const res = await findEzyOrders({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reezyOrders
        dialogFormVisible.value = true
    }
}


// Delete 行
const deleteEzyOrdersFunc = async (row) => {
    const res = await deleteEzyOrders({ ID: row.ID })
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
        quantity: 0,
        seats: '',
        total: 0,
        couponID: 0,
        couponDiscount: 0,
        tax: 0,
        invoicePayment: 0,
        paymentType: '',
        adminDiscount: 0,
        arrivedNumber: 0,
        busNumber: 0,
        createdBy: 0,
        lastEditBy: 0,
        lastActionTime: new Date(),
        source: '',
        notes: '',
        status: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createEzyOrders(formData.value)
                  break
                case 'update':
                  res = await updateEzyOrders(formData.value)
                  break
                default:
                  res = await createEzyOrders(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: 'Create /更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
