<template>
  <div class="container">
    <h1>Edit Docker Container</h1>
    <ContainerForm
        :container="container"
        :availableTags="availableTags"
        submitButtonText="Update Container"
        :onSubmit="submitForm"
    />
  </div>
</template>

<script>
import ContainerForm from "./ContainerForm.vue";
import { ref, onMounted } from "vue";
import { GetAllTypes, GetContainerByID, UpdateContainer } from "../../wailsjs/go/backend/App.js";

export default {
  components: {ContainerForm},
  props: {
    containerId: {
      type: Number,
      required: true,
    },
  },
  setup(props) {
    const container = ref({ name: "", body: "", textWithTags: [] });
    const availableTags = ref([]);

    const fetchTags = async () => {
      const tags = await GetAllTypes();
      availableTags.value = tags.map(tag => ({ id: tag.ID, name: tag.Name }));
    };

    const fetchContainerData = async () => {
      const data = await GetContainerByID(props.containerId);
      container.value = {
        name: data.Name,
        body: data.Body,
        textWithTags: data.public_types.map(item => ({
          text: item.data,
          tag: availableTags.value.find(tag => tag.id === item.type_id)?.name,
          tagId: item.type_id,
        })),
      };
    };

    const submitForm = async () => {
      try {
        const publicTypes = container.value.textWithTags.map(item => ({
          type_id: parseInt(item.tagId),
          data: item.text,
        }));

        const requestData = {
          name: container.value.name,
          body: container.value.body,
          public_types: publicTypes,
        };

        await UpdateContainer(props.containerId, requestData);
        alert(`Container "${container.value.name}" updated successfully!`);
      } catch (error) {
        alert("Error updating container. Please try again.");
        console.error("Failed to update container:", error);
      }
    };

    onMounted(async () => {
      await fetchTags();
      await fetchContainerData();
    });

    return { container, availableTags, submitForm };
  },
};
</script>

<style scoped>
.container {
  padding: 20px;
  max-width: 600px;
  margin: 0 auto;
  background-color: #121212;
  color: #fff;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
}
</style>
