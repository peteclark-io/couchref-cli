'use strict';

import React from 'react';
import CouchRefHeader from '../CouchRefHeader';
import renderer from 'react-test-renderer';

it('Renders ok.', () => {
   const rendered = renderer.create(
      <CouchRefHeader />
   );
   expect(rendered.toJSON()).toMatchSnapshot();
});
