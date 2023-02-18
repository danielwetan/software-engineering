RSpec.describe 'not_to_method' do
  it 'checks that two values do not match' do
    expect(5).not_to eq(6)
    expect('Hello').not_to eq('hello')
    expect([1, 2]).not_to eq([1, 2, 3])
  end
end

# Reference
# https://relishapp.com/rspec/rspec-expectations/v/3-10/docs/built-in-matchers