import React from 'react';
import UserIcon from '@material-ui/icons/Group';
import { Admin, Resource } from 'react-admin';
import jsonServerProvider from 'ra-data-json-server';
import Dashboard from './Dashboard';
import { AttackPatternList,AttackPatternEdit ,AttackPatternCreate } from './AttackPatterns';
import { RelationshipCreate,RelationshipEdit,RelationshipList } from "./Relationship";

const dataProvider = jsonServerProvider("http://localhost:6868");
const Title = () => (<div>Mitre_attack</div>)

const App = () => (
  <Admin title={<Title/>} dashboard={Dashboard} dataProvider={dataProvider}>
      <Resource name="attackPatterns" list={AttackPatternList} edit={AttackPatternEdit}  create={AttackPatternCreate} icon={UserIcon} />
      <Resource name="relationships" list={RelationshipList} edit={RelationshipEdit}  create={RelationshipCreate}  />
  
  </Admin>
)

export default App;
