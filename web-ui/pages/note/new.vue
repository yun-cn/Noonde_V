<template lang="pug">
  #note-new
    c-header-user

    v-content
      .u-pa24
        .u-row-between-middle.u-mb16
          c-breadcrumbs(:items="crumbs")

        .u-row-between-top.u-wrap.u-row-reverse.u-mb16
          c-btn.p-btn.u-ml8() {{ $t('pages.note.new.save-btn') }}
          c-btn.p-btn.u-ml8() {{ $t('pages.note.new.open-btn') }}

        .u-row-between-top.u-wrap.u-row-reverse.u-mb16
          c-text-field.u-mb16.u-full(
            :disabled="loading"
            :label="$t('pages.note.new.input-title')"
            v-model="name"
          )

        .u-row-between-top.u-wrap.u-mb16
          no-ssr
            section(class="container")
              div(
                class="quill-editor"
                :content="content"
                @change="onEditorChange($event)"
                @blur="onEditorBlur($event)"
                @focus="onEditorFocus($event)"
                @ready="onEditorReady($event)"
                v-quill:myQuillEditor="editorOption")
              input(type="file" @change="imageHandler" id="file" hidden)
    c-footer

</template>

<!-- ============================================================================ -->

<script>

  export default {
    data () {
      return {
        crumbs: [
          { label: 'Note', to: '/note'       },
          { label: 'Create', to: '/note/new' },
        ],
        name: '',
        loading:   false,
        content: '<p>I am Example</p>',
        editorOption: {
          // some quill options
          placeholder: this.$t('pages.note.new.input-body'),
          modules: {
            imageResize: true,
            imageResize: {
              displaySize: true
            },
            toolbar: [
              ['bold', 'italic', 'underline', 'strike'],
              ['blockquote', 'code-block'],
              [{ 'header': 1 }, { 'header': 2 }],
              [{ 'list': 'ordered' }, { 'list': 'bullet' }],
              [{ 'align': [] }],
              ['link', 'image', 'video'],
            ]
          }
        }
      }
    },

    methods: {
      onEditorReady(editor) {
        console.log('editor ready!', editor)
      },

      onEditorBlur(editor) {
        console.log('editor blur!', editor)
      },

      onEditorFocus(editor) {
        console.log('editor focus!', editor)
      },

      onEditorChange({ editor, html, text }) {
        console.log('editor change!', editor, html, text);
        this.content = html
      },

      imageHandler(image, callback) {
        var data = new FormData();
        data.append('image', image);
        console.log(image);
        console.log(callback);
        console.log('yayaya');
      },
    },
  }
</script>

<!-- ============================================================================ -->

<style lang="stylus">
  #note-new
    .p-btn
      margin-top     p(4, 0)
      padding-bottom p(0, 8)
    .container
      margin 0 auto
      padding 50px 0
      margin-top -50px
    .quill-editor
      min-height 600px
      max-height 1200px
      overflow-y auto
    .ql-editing
      left 0px !important
      top 0px  !important
</style>
