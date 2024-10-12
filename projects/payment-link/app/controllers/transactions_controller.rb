class TransactionsController < ApplicationController
  def index
    @transactions = Transaction.all
  end

  def new
    @transaction = Transaction.new  # Initialize a new transaction
    @products = Product.all
  end

  def create
    @transaction = Transaction.new(transaction_params)
    product = Product.find(transaction_params[:product_id])
    @transaction.amount = product.price # You can adjust this calculation based on your logic

    if @transaction.save
      redirect_to transactions_path, notice: "Transaction was successfully created."
    else
      render :new
    end
  end

  private

  def transaction_params
    params.require(:transaction).permit(:product_id)
  end
end
