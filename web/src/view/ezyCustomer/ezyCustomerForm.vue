<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="Name:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Phone:" prop="phone">
          <el-input v-model="formData.phone" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Email:" prop="email">
          <el-input v-model="formData.email" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Membership:" prop="membership">
          <el-input v-model="formData.membership" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="Discount:" prop="discount">
          <el-input v-model="formData.discount" :clearable="true" placeholder="请输入" />
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
  name: 'EzyCustomer'
}
</script>

<script setup>
import {
  createEzyCustomer,
  updateEzyCustomer,
  findEzyCustomer
} from '@/api/ezyCustomer'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            name: '',
            phone: '',
            email: '',
            membership: '',
            discount: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               phone : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               email : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEzyCustomer({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reezyCustomer
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
               res = await createEzyCustomer(formData.value)
               break
             case 'update':
               res = await updateEzyCustomer(formData.value)
               break
             default:
               res = await createEzyCustomer(formData.value)
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
