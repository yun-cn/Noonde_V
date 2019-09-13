import Vue from 'vue';

import CBtn         from '@/components/c-btn';
import CTextField   from '@/components/c-text-field';
import CHeader      from '@/components/c-header.vue';
import CHeaderUser  from '@/components/c-header-user.vue';
import CCheckboxes  from '@/components/c-checkboxes';
import CFooter      from '@/components/c-footer.vue';
import CBars        from '@/components/c-bars';
import CBreadcrumbs from '@/components/c-breadcrumbs';
import CChip        from '@/components/c-chip';
import CSearch      from '@/components/c-search';

import {
  CCard,
  CCardTitle,
  CCardText,
  CCardActions,
  CCardDivider,
} from '@/components/c-card/index'



Vue.component('CCard',        CCard);
Vue.component('CBtn',         CBtn);
Vue.component('CHeader',      CHeader);
Vue.component('CHeaderUser',  CHeaderUser);
Vue.component('CCardTitle',   CCardTitle);
Vue.component('CCardText',    CCardText);
Vue.component('CCardActions', CCardActions);
Vue.component('CCardDivider', CCardDivider);
Vue.component('CTextField',   CTextField);
Vue.component('CCheckboxes',  CCheckboxes);
Vue.component('CFooter',      CFooter);
Vue.component('CBars',        CBars);
Vue.component('CBreadcrumbs', CBreadcrumbs);
Vue.component('CChip',        CChip);
Vue.component('CSearch',      CSearch);
