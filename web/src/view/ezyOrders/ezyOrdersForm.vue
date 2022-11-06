<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="Qty:" prop="quantity">
          <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Seats:" prop="seats">
          <el-input v-model="formData.seats" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Total:" prop="total">
          <el-input-number v-model="formData.total" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="Coupon ID:" prop="couponID">
          <el-input v-model.number="formData.couponID" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Discount:" prop="couponDiscount">
          <el-input-number v-model="formData.couponDiscount" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="Tax:" prop="tax">
          <el-input-number v-model="formData.tax" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="Invoice Payment:" prop="invoicePayment">
          <el-input-number v-model="formData.invoicePayment" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="Payment Type:" prop="paymentType">
          <el-input v-model="formData.paymentType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Admin Discount:" prop="adminDiscount">
          <el-input-number v-model="formData.adminDiscount" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="Arrived Number:" prop="arrivedNumber">
          <el-input v-model.number="formData.arrivedNumber" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="BusNumber:" prop="busNumber">
          <el-input v-model.number="formData.busNumber" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Created By:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Last Edit By:" prop="lastEditBy">
          <el-input v-model.number="formData.lastEditBy" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Last Action Time:" prop="lastActionTime">
          <el-date-picker v-model="formData.lastActionTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="Source:" prop="source">
          <el-input v-model="formData.source" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Notes:" prop="notes">
          <el-input v-model="formData.notes" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Status:" prop="status">
          <el-input v-model="formData.status" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
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
  updateEzyOrders,
  findEzyOrders
} from '@/api/ezyOrders'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEzyOrders({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reezyOrders
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
