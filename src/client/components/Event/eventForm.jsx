import React from 'react';
import { Field, reduxForm } from 'redux-form';
import { EventField } from './eventformfields.jsx';
import { DatePicker, TimePicker } from 'redux-form-material-ui'

const { titleField, locationField,
  meetTimeField, descriptionField,
  keyWordField } = EventField;

const EventForm = (props) => {
  const required = value => value == null ? 'Required' : undefined
  const { handleSubmit, pristine,
    reset, submitting,
    eventChange, yelp } = props;
  return (
    <form onSubmit={handleSubmit}>
      <Field name='title' component={titleField} />
      <Field name='location' component={locationField} />
      <Field name='business' component={keyWordField} onChangeAction={eventChange} />
      <button type="button" onClick={yelp}>Yelp</button>
      <Field name="when"
        component={DatePicker}
        format={null}
        onChange={(value) => {
          console.log('date changed ', value) // eslint-disable-line no-console
        }}
        hintText="Day of meeting?"
        validate={required} />
      <Field name="at"
        component={TimePicker}
        format={null}
        defaultValue={null} // TimePicker requires an object,
        onChange={(value) => {
          console.log('time changed ', value) // eslint-disable-line no-console
        }}
        hintText="At what time?"
        validate={required} />
      <Field name='description' component={descriptionField} />
      <div>
        <button type="submit" disabled={pristine || submitting}>Post Event</button>
        <button type="button" disabled={pristine || submitting} onClick={reset}>Reset</button>
      </div>
    </form>
  )
}

const validate = (values) => {
  const errors = {};

  if (!values.title) {
    errors.title = 'Please enter an title';
  }

  if (!values.location) {
    errors.location = 'Please enter a location';
  }

  if (!values.meettime) {
    errors.meettime = 'Please enter a time to meet';
  }

  return errors;
}

export default reduxForm({
  form: 'event-form',
  validate
})(EventForm)