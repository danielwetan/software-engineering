# https://relishapp.com/rspec/rspec-core/v/3-8/docs/example-groups/basic-structure-describe-it
RSpec.describe 'Card' do
  it 'has a type ' do
    card = Card.new('Ace of Spades')
    expect(card.type).to(eq('Ace of Spades'))
    expect(card.type).to eq('Ace of Spades')

    expect(1 + 1).to eq(2)
  end
end

# matchers
# https://www.rubydoc.info/gems/rspec-expectations/RSpec/Matchers
