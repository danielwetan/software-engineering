# frozen_string_literal: true

require 'rspec'

describe 'A fulfilled message expectation' do
  it 'passes' do
    dbl = double(name: 'testing')
    expect(dbl).to receive(:foo)
    expect(dbl.name).to eq('testing')
    dbl.foo
  end
end
