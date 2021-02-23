import React from 'react';
import { Admin, Resource } from 'react-admin';
import jsonServerProvider from 'ra-data-json-server';
import { AttackPatternList,AttackPatternEdit ,AttackPatternCreate } from './AttackPatterns';
import { RelationshipCreate,RelationshipEdit,RelationshipList } from "./Relationship";

const dataProvider = jsonServerProvider("http://localhost:6868");
const Title = () => (<div>Mitre Attack</div>)

const App = () => (
  <Admin title={<Title/>} dataProvider={dataProvider}>
      <Resource name="attackPatterns" list={AttackPatternList} edit={AttackPatternEdit}  create={AttackPatternCreate} />
      <Resource name="relationships" list={RelationshipList} edit={RelationshipEdit}  create={RelationshipCreate}  />
  
  </Admin>
)

export default App;
