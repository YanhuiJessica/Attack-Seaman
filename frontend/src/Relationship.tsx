import React from "react";
import {
  Filter,
  List,
  AutocompleteInput,
  Datagrid,
  SelectInput,
  ReferenceInput,
  TextField,
  ReferenceField,
  Edit,
  SimpleForm,
  TextInput,
  Create,
  required,
} from "react-admin";

const RelationShipTypes = [
  { relationship_type: "uses" },
  { relationship_type: "mitigates" },
  { relationship_type: "subtechnique-of" },
  { relationship_type: "revoked-by" },
];

const RelationshipFilter = (props: any) => (
  <Filter {...props}>
    <SelectInput
      label="RelationshipType"
      choices={RelationShipTypes}
      source="relationship_type"
      optionText="relationship_type"
      optionValue="relationship_type"
      defaultValue="subtechnique-of"
    />
  </Filter>
);

export const RelationshipList = (props: any) => {
  return (
    <List {...props} filters={<RelationshipFilter/>} sort={{ field: 'modified', order: 'DESC' }}>
      <Datagrid rowClick="edit">
        <TextField source="id"/>
        <TextField source="relationship_type"/>

        <ReferenceField
          label="source_ref"
          source="source_ref"
          reference="attackPatterns"
        >
          <TextField source="name"/>
        </ReferenceField>

        <ReferenceField
          label="target_ref"
          source="target_ref"
          reference="attackPatterns"
        >
          <TextField source="name"/>
        </ReferenceField>
      </Datagrid>
    </List>
  );
};

const optionRenderer = (choice: any) =>
  `${choice.external_references[0].external_id} (${choice.name})`;
const RelationshipForm = (props: any) => {
  const disabled = !!props.record.id;
  return (
    <SimpleForm {...props}>
      <TextInput source="type"  defaultValue='relationship' disabled  fullWidth/>
      <TextInput source="id" disabled={disabled} fullWidth/>

      <ReferenceInput
        label="source_ref"
        source="source_ref"
        reference="attackPatterns"
        filterToQuery={(searchText: any) => ({ name: searchText })}
        validate={required()}
        fullWidth
      >
        <AutocompleteInput
          source="name"
          optionText={optionRenderer}
          resettable
          fullWidth
        />
      </ReferenceInput>

      <ReferenceInput
        label="target_ref"
        source="target_ref"
        reference="attackPatterns"
        filterToQuery={(searchText: any) => ({ name: searchText })}
        validate={required()}
        fullWidth
      >
        <AutocompleteInput
          source="name"
          optionText={optionRenderer}
          resettable
          fullWidth
        />
      </ReferenceInput>

      <SelectInput source="relationship_type" label="relationship_type"   choices={RelationShipTypes}
      optionText="relationship_type"
      optionValue="relationship_type"
      defaultValue="subtechnique-of" 
      fullWidth />
    </SimpleForm>
  );
};

export const RelationshipEdit = (props: any) => (
  <Edit title="编辑Relationship" {...props}>
    <RelationshipForm />
  </Edit>
);

export const RelationshipCreate = (props: any) => (
  <Create title="新建一个Relationship" {...props}>
    <RelationshipForm />
  </Create>
);
