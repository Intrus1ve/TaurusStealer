import SideBar from './components/SidebarPlugin';
import lang from 'element-ui/lib/locale/lang/en';
import locale from 'element-ui/lib/locale';
locale.use(lang);

import { 
  Loading, 
  Input,
  Table, 
  TableColumn,
  Button, 
  Tooltip, 
  Dialog, 
  Form, 
  FormItem, 
  Notification, 
  Checkbox, 
  Tag, 
  Progress,
  InputNumber,
  Select, 
  Option,
  Switch,
  ColorPicker,
} from 'element-ui';

import 'bootstrap/scss/bootstrap.scss';
import './assets/sass/style.scss';
import './assets/font-awesome-5.13.0/css/fontawesome-all.min.css'
import 'flag-icon-css/css/flag-icon.min.css'
import './assets/montserrat/montserrat.css'

export default {
  install(Vue) {
    Vue.use(SideBar);
    Vue.use(Loading.directive);
    Vue.component(Input.name, Input);
    Vue.component(Table.name, Table);
    Vue.component(TableColumn.name, TableColumn);
    Vue.component(Button.name, Button);
    Vue.component(Tooltip.name, Tooltip);
    Vue.component(Dialog.name, Dialog);
    Vue.component(Form.name, Form);
    Vue.component(FormItem.name, FormItem);
    Vue.prototype.$notify = Notification;
    Vue.component(Checkbox.name, Checkbox);
    Vue.component(Tag.name, Tag);
    Vue.component(Progress.name, Progress);
    Vue.component(InputNumber.name, InputNumber);
    Vue.component(Select.name, Select);
    Vue.component(Option.name, Option);
    Vue.component(Switch.name, Switch);
    Vue.component(ColorPicker.name, ColorPicker);
  }
};
