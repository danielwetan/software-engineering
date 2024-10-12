class CreatePayments < ActiveRecord::Migration[7.2]
  def change
    create_table :payments do |t|
      t.references :transaction, null: false, foreign_key: true

      t.string :method
      t.integer :amount
      t.string :status

      t.timestamps
    end
  end
end
