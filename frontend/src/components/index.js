import FormGroupInput from "./Inputs/formGroupInput.vue";

import DropDown from "./Dropdown.vue";
import PaperTable from "./PaperTable.vue";
import Button from "./Button";
import Overlay from "./Overlay";

import Card from "./Cards/Card.vue";
import ChartCard from "./Cards/ChartCard.vue";
import StatsCard from "./Cards/StatsCard.vue";
import WideCard from "./Cards/WideCard.vue";

import Password from "./Inputs/Password.vue";
import FileSelector from "./Inputs/FileSelector.vue";

import SidebarPlugin from "./SidebarPlugin/index";

let components = {
  FormGroupInput,
  Card,
  ChartCard,
  StatsCard,
  WideCard,
  Password,
  FileSelector,
  PaperTable,
  DropDown,
  SidebarPlugin,
  Overlay
};

export default components;

export {
  FormGroupInput,
  Card,
  ChartCard,
  StatsCard,
  WideCard,
  Password,
  FileSelector,
  PaperTable,
  DropDown,
  Button,
  SidebarPlugin,
  Overlay
};
