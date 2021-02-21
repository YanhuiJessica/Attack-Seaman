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
  BooleanInput,
  Edit,
  SimpleForm,
  TextInput,
  DateInput,
  ArrayInput,
  SimpleFormIterator,
  Create,
} from "react-admin";
import MyUrlField from "./MyUrlField";

const RelationShipTypes = [
  { relationship_type: "uses" },
  { relationship_type: "mitigates" },
  { relationship_type: "subtechnique-of" },
  { relationship_type: "revoked-by" },
];

const RelationshipFilter = (props: any) => (
  <Filter {...props}>
    {/* <TextInput label="Name" source="name" alwaysOn /> */}
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
    <List {...props} filters={<RelationshipFilter />}>
      <Datagrid rowClick="edit">
        <TextField source="id" />
        <TextField source="relationship_type" />

        <ReferenceField
          label="source_ref"
          source="source_ref"
          reference="attackPatterns"
        >
          <TextField source="name" />
        </ReferenceField>

        <ReferenceField
          label="target_ref"
          source="target_ref"
          reference="attackPatterns"
        >
          <TextField source="name" />
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
      <TextInput source="id" disabled={disabled} />
      <TextInput source="type" />

      <ReferenceInput
        label="source_ref"
        source="source_ref"
        reference="attackPatterns"
        filterToQuery={(searchText: any) => ({ name: searchText })}
      >
        <AutocompleteInput
          source="name"
          optionText={optionRenderer}
          resettable
        />
      </ReferenceInput>

      <ReferenceInput
        label="target_ref"
        source="target_ref"
        reference="attackPatterns"
        filterToQuery={(searchText: any) => ({ name: searchText })}
      >
        <AutocompleteInput
          source="name"
          optionText={optionRenderer}
          resettable
        />
      </ReferenceInput>

      <TextInput source="relationship_type" label="relationship_type" />
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
