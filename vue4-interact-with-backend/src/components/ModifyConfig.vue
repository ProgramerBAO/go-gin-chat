<template>
  <div>
    <h1>Vue Form with Axios</h1>

    <!-- 三个信息的表单 -->
    <div>
      <h2>信息</h2>
      <input v-model="formData.newInfo" placeholder="新增信息"/>
      <input v-model="formData.updateInfo" placeholder="修改信息"/>
      <input v-model="formData.originalInfo" placeholder="原始信息"/>
      <button @click="submitData">提交</button>
    </div>

    <!-- 弹窗显示结果 -->
    <div v-if="showResult">
      <h2>结果</h2>
      <p>{{ result }}</p>
      <button class="copy-button" @click="copyResult">复制结果</button>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import ClipboardJS from "clipboard";

export default {
  data() {
    return {
      formData: {
        newInfo: "",
        updateInfo: "",
        originalInfo: ""
      },
      result: "",
      showResult: false,
      clipboard: null,
    };
  },
  mounted() {
    // 初始化ClipboardJS，将要复制的文本设置为 result 的值
    this.clipboard = new ClipboardJS(".copy-button", {
      text: () => this.result,
    });
  },
  methods: {
    async submitData() {
      try {
        const response = await axios.post("api/new", this.formData);
        console.log(this.formData)
        this.result = response.data.message;
        this.showResult = true;
      } catch (error) {
        this.result = "请求失败";
        this.showResult = true;
      }
    },
    copyResult() {
      if (this.clipboard) {
        this.clipboard.on("success", () => {
          // 复制成功后，在这里显示提示弹窗
          window.alert("已成功复制到剪切板！");
        });
      }
    },
  },
};
</script>
