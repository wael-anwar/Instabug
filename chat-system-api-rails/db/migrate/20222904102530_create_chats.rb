class CreateChats < ActiveRecord::Migration[6.0]
  def change
    create_table :chats do |t|
      t.integer :number
      t.references :application, null: false, foreign_key: true,  type: :bigint
      t.integer :message_count
      t.timestamps
    end
  end
end
