import FormGroupInput from "./Inputs/formGroupInput.vue";

import DropDown from "./Dropdown.vue";
import PaperTable from "./PaperTable.vue";
import Button from "./Button";

import Card from "./Cards/Card.vue";
import ChartCard from "./Cards/ChartCard.vue";
import StatsCard from "./Cards/StatsCard.vue";
import WideCard from "./Cards/WideCard.vue";

import SidebarPlugin from "./SidebarPlugin/index";

let components = {
  FormGroupInput,
  Card,
  ChartCard,
  StatsCard,
  WideCard,
  PaperTable,
  DropDown,
  SidebarPlugin
};

export default components;

export {
  FormGroupInput,
  Card,
  ChartCard,
  StatsCard,
  WideCard,
  PaperTable,
  DropDown,
  Button,
  SidebarPlugin
};
