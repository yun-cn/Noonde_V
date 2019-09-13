import Vue from 'vue';
import Quill from 'quill'
import VueQuillEditor from 'vue-quill-editor/dist/ssr';

import { ImageDrop } from 'quill-image-drop-module'
import ImageResize from 'quill-image-resize-module'


import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'


Quill.register('modules/imageDrop', ImageDrop);
Quill.register('modules/imageResize', ImageResize);


Vue.use(VueQuillEditor);
