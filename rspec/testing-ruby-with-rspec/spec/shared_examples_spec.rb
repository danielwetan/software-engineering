RSpec.shared_examples 'a Ruby object with three elements' do 
  it 'returns the number of items' do 
    expect(subject.length).to eq(3)
  end
end

RSpec.describe Array do 
  subject { [1, 2, 3] }
  include_examples 'a Ruby object with three elements'
end

RSpec.describe String do 
  subject { 'abc' }
  include_examples 'a Ruby object with three elements'
end

RSpec.describe Hash do 
  subject{ {a: 1, b: 2, c: 3}  }
  include_examples 'a Ruby object with three elements'
end

class SausageLink 
  def length 
    3
  end
end

RSpec.describe SausageLink do 
  subject { described_class.new }
end

# RSpec shared examples reference
# https://relishapp.com/rspec/rspec-core/v/3-12/docs/example-groups/shared-examples

